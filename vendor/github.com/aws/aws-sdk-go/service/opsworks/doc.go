// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package opsworks provides the client and types for making API
// requests to AWS OpsWorks.
//
// Welcome to the AWS OpsWorks Stacks API Reference. This guide provides descriptions,
// syntax, and usage examples for AWS OpsWorks Stacks actions and data types,
// including common parameters and error codes.
//
// AWS OpsWorks Stacks is an application management service that provides an
// integrated experience for overseeing the complete application lifecycle.
// For information about this product, go to the AWS OpsWorks (http://aws.amazon.com/opsworks/)
// details page.
//
// SDKs and CLI
//
// The most common way to use the AWS OpsWorks Stacks API is by using the AWS
// Command Line Interface (CLI) or by using one of the AWS SDKs to implement
// applications in your preferred language. For more information, see:
//
//    * AWS CLI (http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html)
//
//    * AWS SDK for Java (http://docs.aws.amazon.com/AWSJavaSDK/latest/javadoc/com/amazonaws/services/opsworks/AWSOpsWorksClient.html)
//
//    * AWS SDK for .NET (http://docs.aws.amazon.com/sdkfornet/latest/apidocs/html/N_Amazon_OpsWorks.htm)
//
//    * AWS SDK for PHP 2 (http://docs.aws.amazon.com/aws-sdk-php-2/latest/class-Aws.OpsWorks.OpsWorksClient.html)
//
//    * AWS SDK for Ruby (http://docs.aws.amazon.com/sdkforruby/api/)
//
//    * AWS SDK for Node.js (http://aws.amazon.com/documentation/sdkforjavascript/)
//
//    * AWS SDK for Python(Boto) (http://docs.pythonboto.org/en/latest/ref/opsworks.html)
//
// Endpoints
//
// AWS OpsWorks Stacks supports the following endpoints, all HTTPS. You must
// connect to one of the following endpoints. Stacks can only be accessed or
// managed within the endpoint in which they are created.
//
//    * opsworks.us-east-1.amazonaws.com
//
//    * opsworks.us-east-2.amazonaws.com
//
//    * opsworks.us-west-1.amazonaws.com
//
//    * opsworks.us-west-2.amazonaws.com
//
//    * opsworks.eu-west-1.amazonaws.com
//
//    * opsworks.eu-west-2.amazonaws.com
//
//    * opsworks.eu-central-1.amazonaws.com
//
//    * opsworks.ap-northeast-1.amazonaws.com
//
//    * opsworks.ap-northeast-2.amazonaws.com
//
//    * opsworks.ap-south-1.amazonaws.com
//
//    * opsworks.ap-southeast-1.amazonaws.com
//
//    * opsworks.ap-southeast-2.amazonaws.com
//
//    * opsworks.sa-east-1.amazonaws.com
//
// Chef Versions
//
// When you call CreateStack, CloneStack, or UpdateStack we recommend you use
// the ConfigurationManager parameter to specify the Chef version. The recommended
// and default value for Linux stacks is currently 12. Windows stacks use Chef
// 12.2. For more information, see Chef Versions (http://docs.aws.amazon.com/opsworks/latest/userguide/workingcookbook-chef11.html).
//
// You can specify Chef 12, 11.10, or 11.4 for your Linux stack. We recommend
// migrating your existing Linux stacks to Chef 12 as soon as possible.
//
// See https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18 for more information on this service.
//
// See opsworks package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/opsworks/
//
// Using the Client
//
// To use the client for AWS OpsWorks you will first need
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
//   svc := opsworks.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS OpsWorks client OpsWorks for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/opsworks/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.AssignInstance(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("AssignInstance result:")
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
//   result, err := svc.AssignInstanceWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package opsworks
