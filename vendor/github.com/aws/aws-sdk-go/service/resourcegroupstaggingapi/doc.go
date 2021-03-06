// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package resourcegroupstaggingapi provides the client and types for making API
// requests to AWS Resource Groups Tagging API.
//
// This guide describes the API operations for the resource groups tagging.
//
// A tag is a label that you assign to an AWS resource. A tag consists of a
// key and a value, both of which you define. For example, if you have two Amazon
// EC2 instances, you might assign both a tag key of "Stack." But the value
// of "Stack" might be "Testing" for one and "Production" for the other.
//
// Tagging can help you organize your resources and enables you to simplify
// resource management, access management and cost allocation. For more information
// about tagging, see Working with Tag Editor (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/tag-editor.html)
// and Working with Resource Groups (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/resource-groups.html).
// For more information about permissions you need to use the resource groups
// tagging APIs, see Obtaining Permissions for Resource Groups  (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/obtaining-permissions-for-resource-groups.html)
// and Obtaining Permissions for Tagging  (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/obtaining-permissions-for-tagging.html).
//
// You can use the resource groups tagging APIs to complete the following tasks:
//
//    * Tag and untag supported resources located in the specified region for
//    the AWS account
//
//    * Use tag-based filters to search for resources located in the specified
//    region for the AWS account
//
//    * List all existing tag keys in the specified region for the AWS account
//
//    * List all existing values for the specified key in the specified region
//    for the AWS account
//
// Not all resources can have tags. For a lists of resources that you can tag,
// see Supported Resources (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/supported-resources.html)
// in the AWS Resource Groups and Tag Editor User Guide.
//
// To make full use of the resource groups tagging APIs, you might need additional
// IAM permissions, including permission to access the resources of individual
// services as well as permission to view and apply tags to those resources.
// For more information, see Obtaining Permissions for Tagging (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/obtaining-permissions-for-tagging.html)
// in the AWS Resource Groups and Tag Editor User Guide.
//
// See https://docs.aws.amazon.com/goto/WebAPI/resourcegroupstaggingapi-2017-01-26 for more information on this service.
//
// See resourcegroupstaggingapi package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/resourcegroupstaggingapi/
//
// Using the Client
//
// To use the client for AWS Resource Groups Tagging API you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := resourcegroupstaggingapi.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Resource Groups Tagging API client ResourceGroupsTaggingAPI for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/resourcegroupstaggingapi/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.GetResources(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("GetResources result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.GetResourcesWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package resourcegroupstaggingapi
