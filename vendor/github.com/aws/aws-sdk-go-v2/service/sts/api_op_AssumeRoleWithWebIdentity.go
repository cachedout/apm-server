// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/AssumeRoleWithWebIdentityRequest
type AssumeRoleWithWebIdentityInput struct {
	_ struct{} `type:"structure"`

	// The duration, in seconds, of the role session. The value can range from 900
	// seconds (15 minutes) up to the maximum session duration setting for the role.
	// This setting can have a value from 1 hour to 12 hours. If you specify a value
	// higher than this setting, the operation fails. For example, if you specify
	// a session duration of 12 hours, but your administrator set the maximum session
	// duration to 6 hours, your operation fails. To learn how to view the maximum
	// value for your role, see View the Maximum Session Duration Setting for a
	// Role (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html#id_roles_use_view-role-max-session)
	// in the IAM User Guide.
	//
	// By default, the value is set to 3600 seconds.
	//
	// The DurationSeconds parameter is separate from the duration of a console
	// session that you might request using the returned credentials. The request
	// to the federation endpoint for a console sign-in token takes a SessionDuration
	// parameter that specifies the maximum length of the console session. For more
	// information, see Creating a URL that Enables Federated Users to Access the
	// AWS Management Console (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-custom-url.html)
	// in the IAM User Guide.
	DurationSeconds *int64 `min:"900" type:"integer"`

	// An IAM policy in JSON format that you want to use as an inline session policy.
	//
	// This parameter is optional. Passing policies to this operation returns new
	// temporary credentials. The resulting session's permissions are the intersection
	// of the role's identity-based policy and the session policies. You can use
	// the role's temporary credentials in subsequent AWS API calls to access resources
	// in the account that owns the role. You cannot use session policies to grant
	// more permissions than those allowed by the identity-based policy of the role
	// that is being assumed. For more information, see Session Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// in the IAM User Guide.
	//
	// The plain text that you use for both inline and managed session policies
	// shouldn't exceed 2048 characters. The JSON policy characters can be any ASCII
	// character from the space character to the end of the valid character list
	// (\u0020 through \u00FF). It can also include the tab (\u0009), linefeed (\u000A),
	// and carriage return (\u000D) characters.
	//
	// The characters in this parameter count towards the 2048 character session
	// policy guideline. However, an AWS conversion compresses the session policies
	// into a packed binary format that has a separate limit. This is the enforced
	// limit. The PackedPolicySize response element indicates by percentage how
	// close the policy is to the upper size limit.
	Policy *string `min:"1" type:"string"`

	// The Amazon Resource Names (ARNs) of the IAM managed policies that you want
	// to use as managed session policies. The policies must exist in the same account
	// as the role.
	//
	// This parameter is optional. You can provide up to 10 managed policy ARNs.
	// However, the plain text that you use for both inline and managed session
	// policies shouldn't exceed 2048 characters. For more information about ARNs,
	// see Amazon Resource Names (ARNs) and AWS Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// The characters in this parameter count towards the 2048 character session
	// policy guideline. However, an AWS conversion compresses the session policies
	// into a packed binary format that has a separate limit. This is the enforced
	// limit. The PackedPolicySize response element indicates by percentage how
	// close the policy is to the upper size limit.
	//
	// Passing policies to this operation returns new temporary credentials. The
	// resulting session's permissions are the intersection of the role's identity-based
	// policy and the session policies. You can use the role's temporary credentials
	// in subsequent AWS API calls to access resources in the account that owns
	// the role. You cannot use session policies to grant more permissions than
	// those allowed by the identity-based policy of the role that is being assumed.
	// For more information, see Session Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
	// in the IAM User Guide.
	PolicyArns []PolicyDescriptorType `type:"list"`

	// The fully qualified host component of the domain name of the identity provider.
	//
	// Specify this value only for OAuth 2.0 access tokens. Currently www.amazon.com
	// and graph.facebook.com are the only supported identity providers for OAuth
	// 2.0 access tokens. Do not include URL schemes and port numbers.
	//
	// Do not specify this value for OpenID Connect ID tokens.
	ProviderId *string `min:"4" type:"string"`

	// The Amazon Resource Name (ARN) of the role that the caller is assuming.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`

	// An identifier for the assumed role session. Typically, you pass the name
	// or identifier that is associated with the user who is using your application.
	// That way, the temporary security credentials that your application will use
	// are associated with that user. This session name is included as part of the
	// ARN and assumed role ID in the AssumedRoleUser response element.
	//
	// The regex used to validate this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can
	// also include underscores or any of the following characters: =,.@-
	//
	// RoleSessionName is a required field
	RoleSessionName *string `min:"2" type:"string" required:"true"`

	// The OAuth 2.0 access token or OpenID Connect ID token that is provided by
	// the identity provider. Your application must get this token by authenticating
	// the user who is using your application with a web identity provider before
	// the application makes an AssumeRoleWithWebIdentity call.
	//
	// WebIdentityToken is a required field
	WebIdentityToken *string `min:"4" type:"string" required:"true"`
}

// String returns the string representation
func (s AssumeRoleWithWebIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssumeRoleWithWebIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssumeRoleWithWebIdentityInput"}
	if s.DurationSeconds != nil && *s.DurationSeconds < 900 {
		invalidParams.Add(aws.NewErrParamMinValue("DurationSeconds", 900))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Policy", 1))
	}
	if s.ProviderId != nil && len(*s.ProviderId) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("ProviderId", 4))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}

	if s.RoleSessionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleSessionName"))
	}
	if s.RoleSessionName != nil && len(*s.RoleSessionName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleSessionName", 2))
	}

	if s.WebIdentityToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebIdentityToken"))
	}
	if s.WebIdentityToken != nil && len(*s.WebIdentityToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("WebIdentityToken", 4))
	}
	if s.PolicyArns != nil {
		for i, v := range s.PolicyArns {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PolicyArns", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful AssumeRoleWithWebIdentity request,
// including temporary AWS credentials that can be used to make AWS requests.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/AssumeRoleWithWebIdentityResponse
type AssumeRoleWithWebIdentityOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) and the assumed role ID, which are identifiers
	// that you can use to refer to the resulting temporary security credentials.
	// For example, you can reference these credentials as a principal in a resource-based
	// policy by using the ARN or assumed role ID. The ARN and ID include the RoleSessionName
	// that you specified when you called AssumeRole.
	AssumedRoleUser *AssumedRoleUser `type:"structure"`

	// The intended audience (also known as client ID) of the web identity token.
	// This is traditionally the client identifier issued to the application that
	// requested the web identity token.
	Audience *string `type:"string"`

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security token.
	//
	// The size of the security token that STS API operations return is not fixed.
	// We strongly recommend that you make no assumptions about the maximum size.
	Credentials *Credentials `type:"structure"`

	// A percentage value that indicates the size of the policy in packed form.
	// The service rejects any policy with a packed size greater than 100 percent,
	// which means the policy exceeded the allowed space.
	PackedPolicySize *int64 `type:"integer"`

	// The issuing authority of the web identity token presented. For OpenID Connect
	// ID tokens, this contains the value of the iss field. For OAuth 2.0 access
	// tokens, this contains the value of the ProviderId parameter that was passed
	// in the AssumeRoleWithWebIdentity request.
	Provider *string `type:"string"`

	// The unique user identifier that is returned by the identity provider. This
	// identifier is associated with the WebIdentityToken that was submitted with
	// the AssumeRoleWithWebIdentity call. The identifier is typically unique to
	// the user and the application that acquired the WebIdentityToken (pairwise
	// identifier). For OpenID Connect ID tokens, this field contains the value
	// returned by the identity provider as the token's sub (Subject) claim.
	SubjectFromWebIdentityToken *string `min:"6" type:"string"`
}

// String returns the string representation
func (s AssumeRoleWithWebIdentityOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssumeRoleWithWebIdentity = "AssumeRoleWithWebIdentity"

// AssumeRoleWithWebIdentityRequest returns a request value for making API operation for
// AWS Security Token Service.
//
// Returns a set of temporary security credentials for users who have been authenticated
// in a mobile or web application with a web identity provider. Example providers
// include Amazon Cognito, Login with Amazon, Facebook, Google, or any OpenID
// Connect-compatible identity provider.
//
// For mobile applications, we recommend that you use Amazon Cognito. You can
// use Amazon Cognito with the AWS SDK for iOS Developer Guide (http://aws.amazon.com/sdkforios/)
// and the AWS SDK for Android Developer Guide (http://aws.amazon.com/sdkforandroid/)
// to uniquely identify a user. You can also supply the user with a consistent
// identity throughout the lifetime of an application.
//
// To learn more about Amazon Cognito, see Amazon Cognito Overview (https://docs.aws.amazon.com/mobile/sdkforandroid/developerguide/cognito-auth.html#d0e840)
// in AWS SDK for Android Developer Guide and Amazon Cognito Overview (https://docs.aws.amazon.com/mobile/sdkforios/developerguide/cognito-auth.html#d0e664)
// in the AWS SDK for iOS Developer Guide.
//
// Calling AssumeRoleWithWebIdentity does not require the use of AWS security
// credentials. Therefore, you can distribute an application (for example, on
// mobile devices) that requests temporary security credentials without including
// long-term AWS credentials in the application. You also don't need to deploy
// server-based proxy services that use long-term AWS credentials. Instead,
// the identity of the caller is validated by using a token from the web identity
// provider. For a comparison of AssumeRoleWithWebIdentity with the other API
// operations that produce temporary credentials, see Requesting Temporary Security
// Credentials (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS API operations (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide.
//
// The temporary security credentials returned by this API consist of an access
// key ID, a secret access key, and a security token. Applications can use these
// temporary security credentials to sign calls to AWS service API operations.
//
// By default, the temporary security credentials created by AssumeRoleWithWebIdentity
// last for one hour. However, you can use the optional DurationSeconds parameter
// to specify the duration of your session. You can provide a value from 900
// seconds (15 minutes) up to the maximum session duration setting for the role.
// This setting can have a value from 1 hour to 12 hours. To learn how to view
// the maximum value for your role, see View the Maximum Session Duration Setting
// for a Role (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html#id_roles_use_view-role-max-session)
// in the IAM User Guide. The maximum session duration limit applies when you
// use the AssumeRole* API operations or the assume-role* CLI commands. However
// the limit does not apply when you use those operations to create a console
// URL. For more information, see Using IAM Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html)
// in the IAM User Guide.
//
// The temporary security credentials created by AssumeRoleWithWebIdentity can
// be used to make API calls to any AWS service with the following exception:
// you cannot call the STS GetFederationToken or GetSessionToken API operations.
//
// (Optional) You can pass inline or managed session policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// to this operation. You can pass a single JSON policy document to use as an
// inline session policy. You can also specify up to 10 managed policies to
// use as managed session policies. The plain text that you use for both inline
// and managed session policies shouldn't exceed 2048 characters. Passing policies
// to this operation returns new temporary credentials. The resulting session's
// permissions are the intersection of the role's identity-based policy and
// the session policies. You can use the role's temporary credentials in subsequent
// AWS API calls to access resources in the account that owns the role. You
// cannot use session policies to grant more permissions than those allowed
// by the identity-based policy of the role that is being assumed. For more
// information, see Session Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// in the IAM User Guide.
//
// Before your application can call AssumeRoleWithWebIdentity, you must have
// an identity token from a supported identity provider and create a role that
// the application can assume. The role that your application assumes must trust
// the identity provider that is associated with the identity token. In other
// words, the identity provider must be specified in the role's trust policy.
//
// Calling AssumeRoleWithWebIdentity can result in an entry in your AWS CloudTrail
// logs. The entry includes the Subject (http://openid.net/specs/openid-connect-core-1_0.html#Claims)
// of the provided Web Identity Token. We recommend that you avoid using any
// personally identifiable information (PII) in this field. For example, you
// could instead use a GUID or a pairwise identifier, as suggested in the OIDC
// specification (http://openid.net/specs/openid-connect-core-1_0.html#SubjectIDTypes).
//
// For more information about how to use web identity federation and the AssumeRoleWithWebIdentity
// API, see the following resources:
//
//    * Using Web Identity Federation API Operations for Mobile Apps (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc_manual.html)
//    and Federation Through a Web-based Identity Provider (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_assumerolewithwebidentity).
//
//    * Web Identity Federation Playground (https://web-identity-federation-playground.s3.amazonaws.com/index.html).
//    Walk through the process of authenticating through Login with Amazon,
//    Facebook, or Google, getting temporary security credentials, and then
//    using those credentials to make a request to AWS.
//
//    * AWS SDK for iOS Developer Guide (http://aws.amazon.com/sdkforios/) and
//    AWS SDK for Android Developer Guide (http://aws.amazon.com/sdkforandroid/).
//    These toolkits contain sample apps that show how to invoke the identity
//    providers, and then how to use the information from these providers to
//    get and use temporary security credentials.
//
//    * Web Identity Federation with Mobile Applications (http://aws.amazon.com/articles/web-identity-federation-with-mobile-applications).
//    This article discusses web identity federation and shows an example of
//    how to use web identity federation to get access to content in Amazon
//    S3.
//
//    // Example sending a request using AssumeRoleWithWebIdentityRequest.
//    req := client.AssumeRoleWithWebIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/AssumeRoleWithWebIdentity
func (c *Client) AssumeRoleWithWebIdentityRequest(input *AssumeRoleWithWebIdentityInput) AssumeRoleWithWebIdentityRequest {
	op := &aws.Operation{
		Name:       opAssumeRoleWithWebIdentity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssumeRoleWithWebIdentityInput{}
	}

	req := c.newRequest(op, input, &AssumeRoleWithWebIdentityOutput{})
	req.Config.Credentials = aws.AnonymousCredentials
	return AssumeRoleWithWebIdentityRequest{Request: req, Input: input, Copy: c.AssumeRoleWithWebIdentityRequest}
}

// AssumeRoleWithWebIdentityRequest is the request type for the
// AssumeRoleWithWebIdentity API operation.
type AssumeRoleWithWebIdentityRequest struct {
	*aws.Request
	Input *AssumeRoleWithWebIdentityInput
	Copy  func(*AssumeRoleWithWebIdentityInput) AssumeRoleWithWebIdentityRequest
}

// Send marshals and sends the AssumeRoleWithWebIdentity API request.
func (r AssumeRoleWithWebIdentityRequest) Send(ctx context.Context) (*AssumeRoleWithWebIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssumeRoleWithWebIdentityResponse{
		AssumeRoleWithWebIdentityOutput: r.Request.Data.(*AssumeRoleWithWebIdentityOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssumeRoleWithWebIdentityResponse is the response type for the
// AssumeRoleWithWebIdentity API operation.
type AssumeRoleWithWebIdentityResponse struct {
	*AssumeRoleWithWebIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssumeRoleWithWebIdentity request.
func (r *AssumeRoleWithWebIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
