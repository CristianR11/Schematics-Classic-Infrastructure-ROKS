package ibm

import (
	"fmt"
	"strings"

	apigatewaysdk "github.com/IBM/apigateway-go-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIBMApiGatewayEndpointSubscription() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMApiGatewayEndpointSubscriptionCreate,
		Read:     resourceIBMApiGatewayEndpointSubscriptionGet,
		Update:   resourceIBMApiGatewayEndpointSubscriptionUpdate,
		Delete:   resourceIBMApiGatewayEndpointSubscriptionDelete,
		Importer: &schema.ResourceImporter{},
		Exists:   resourceIBMApiGatewayEndpointSubscriptionExists,
		Schema: map[string]*schema.Schema{
			"artifact_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Endpoint ID",
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subscription Id",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subscription name",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subscription type. Allowable values are external, bluemix",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Client Sercret of a Subscription",
			},
			"secret_provided": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if client secret is provided to subscription or not",
			},
		},
	}
}
func resourceIBMApiGatewayEndpointSubscriptionCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}
	endpointservice, err := meta.(ClientSession).APIGateway()
	if err != nil {
		return err
	}
	payload := &apigatewaysdk.CreateSubscriptionOptions{}

	oauthtoken := sess.Config.IAMAccessToken
	oauthtoken = strings.Replace(oauthtoken, "Bearer ", "", -1)
	payload.Authorization = &oauthtoken

	artifactID := d.Get("artifact_id").(string)
	payload.ArtifactID = &artifactID

	clientID := d.Get("client_id").(string)
	payload.ClientID = &clientID

	var name string
	if v, ok := d.GetOk("name"); ok && v != nil {
		name = v.(string)
		payload.Name = &name
	}
	var shareType string
	if v, ok := d.GetOk("type"); ok && v != nil {
		shareType = v.(string)
		payload.Type = &shareType
	}
	var clientSecret string
	if v, ok := d.GetOk("client_secret"); ok && v != nil {
		clientSecret = v.(string)
		payload.ClientSecret = &clientSecret
	}

	result, response, err := endpointservice.CreateSubscription(payload)
	if err != nil {
		return fmt.Errorf("Error creating Subscription: %s %d", err, response.StatusCode)
	}
	d.SetId(fmt.Sprintf("%s//%s", *result.ArtifactID, *result.ClientID))

	return resourceIBMApiGatewayEndpointSubscriptionGet(d, meta)
}

func resourceIBMApiGatewayEndpointSubscriptionGet(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}
	endpointservice, err := meta.(ClientSession).APIGateway()
	if err != nil {
		return err
	}

	parts := d.Id()
	partslist := strings.Split(parts, "//")
	artifactID := partslist[0]
	clientID := partslist[1]

	oauthtoken := sess.Config.IAMAccessToken
	oauthtoken = strings.Replace(oauthtoken, "Bearer ", "", -1)

	payload := apigatewaysdk.GetSubscriptionOptions{
		ArtifactID:    &artifactID,
		ID:            &clientID,
		Authorization: &oauthtoken,
	}
	result, response, err := endpointservice.GetSubscription(&payload)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting Subscription: %s", err)
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	d.Set("artifact_id", result.ArtifactID)
	d.Set("client_id", result.ClientID)
	d.Set("type", result.Type)
	if result.Name != nil {
		d.Set("name", result.Name)
	}
	if v, ok := d.GetOk("client_secret"); ok && v != nil {
		*result.SecretProvided = true
		d.Set("secret_provided", result.SecretProvided)
	}
	return nil
}

func resourceIBMApiGatewayEndpointSubscriptionUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}
	endpointservice, err := meta.(ClientSession).APIGateway()
	if err != nil {
		return err
	}
	payload := &apigatewaysdk.UpdateSubscriptionOptions{}

	parts := d.Id()
	partslist := strings.Split(parts, "//")
	artifactID := partslist[0]
	clientID := partslist[1]

	oauthtoken := sess.Config.IAMAccessToken
	oauthtoken = strings.Replace(oauthtoken, "Bearer ", "", -1)
	payload.Authorization = &oauthtoken

	payload.ID = &clientID
	payload.NewClientID = &clientID

	payload.ArtifactID = &artifactID
	payload.NewArtifactID = &artifactID

	name := d.Get("name").(string)
	payload.NewName = &name

	update := false

	if d.HasChange("name") {
		name := d.Get("name").(string)
		payload.NewName = &name
		update = true
	}
	if d.HasChange("client_secret") {
		clientSecret := d.Get("client_secret").(string)
		secretpayload := &apigatewaysdk.AddSubscriptionSecretOptions{
			Authorization: &oauthtoken,
			ArtifactID:    &artifactID,
			ID:            &clientID,
			ClientSecret:  &clientSecret,
		}
		_, SecretResponse, err := endpointservice.AddSubscriptionSecret(secretpayload)
		if err != nil {
			return fmt.Errorf("Error Adding Secret to Subscription: %s,%d", err, SecretResponse.StatusCode)
		}
	}
	if update {
		_, response, err := endpointservice.UpdateSubscription(payload)
		if err != nil {
			return fmt.Errorf("Error updating Subscription: %s,%d", err, response.StatusCode)
		}
	}
	return resourceIBMApiGatewayEndpointSubscriptionGet(d, meta)
}

func resourceIBMApiGatewayEndpointSubscriptionDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}
	endpointservice, err := meta.(ClientSession).APIGateway()
	if err != nil {
		return err
	}
	parts := d.Id()
	partslist := strings.Split(parts, "//")
	artifactID := partslist[0]
	clientID := partslist[1]

	oauthtoken := sess.Config.IAMAccessToken
	oauthtoken = strings.Replace(oauthtoken, "Bearer ", "", -1)

	payload := apigatewaysdk.DeleteSubscriptionOptions{
		ArtifactID:    &artifactID,
		ID:            &clientID,
		Authorization: &oauthtoken,
	}
	response, err := endpointservice.DeleteSubscription(&payload)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error deleting Subscription: %s", err)
	}
	d.SetId("")

	return nil
}

func resourceIBMApiGatewayEndpointSubscriptionExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return false, err
	}
	endpointservice, err := meta.(ClientSession).APIGateway()
	if err != nil {
		return false, err
	}
	parts := d.Id()
	partslist := strings.Split(parts, "//")
	artifactID := partslist[0]
	clientID := partslist[1]

	oauthtoken := sess.Config.IAMAccessToken
	oauthtoken = strings.Replace(oauthtoken, "Bearer ", "", -1)

	payload := apigatewaysdk.GetSubscriptionOptions{
		ArtifactID:    &artifactID,
		ID:            &clientID,
		Authorization: &oauthtoken,
	}
	_, response, err := endpointservice.GetSubscription(&payload)
	if err != nil && response.StatusCode != 404 {
		if response.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
