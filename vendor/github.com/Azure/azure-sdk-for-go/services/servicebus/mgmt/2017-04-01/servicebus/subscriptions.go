package servicebus

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// SubscriptionsClient is the azure Service Bus client
type SubscriptionsClient struct {
	BaseClient
}

// NewSubscriptionsClient creates an instance of the SubscriptionsClient client.
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
	return NewSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSubscriptionsClientWithBaseURI creates an instance of the SubscriptionsClient client.
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
	return SubscriptionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a topic subscription.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
// topicName - the topic name.
// subscriptionName - the subscription name.
// parameters - parameters supplied to create a subscription resource.
func (client SubscriptionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, parameters SBSubscription) (result SBSubscription, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: topicName,
			Constraints: []validation.Constraint{{Target: "topicName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: subscriptionName,
			Constraints: []validation.Constraint{{Target: "subscriptionName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "subscriptionName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicebus.SubscriptionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, namespaceName, topicName, subscriptionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SubscriptionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, parameters SBSubscription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"subscriptionName":  autorest.Encode("path", subscriptionName),
		"topicName":         autorest.Encode("path", topicName),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) CreateOrUpdateResponder(resp *http.Response) (result SBSubscription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a subscription from the specified topic.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
// topicName - the topic name.
// subscriptionName - the subscription name.
func (client SubscriptionsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: topicName,
			Constraints: []validation.Constraint{{Target: "topicName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: subscriptionName,
			Constraints: []validation.Constraint{{Target: "subscriptionName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "subscriptionName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicebus.SubscriptionsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, namespaceName, topicName, subscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SubscriptionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"subscriptionName":  autorest.Encode("path", subscriptionName),
		"topicName":         autorest.Encode("path", topicName),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns a subscription description for the specified topic.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
// topicName - the topic name.
// subscriptionName - the subscription name.
func (client SubscriptionsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string) (result SBSubscription, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: topicName,
			Constraints: []validation.Constraint{{Target: "topicName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: subscriptionName,
			Constraints: []validation.Constraint{{Target: "subscriptionName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "subscriptionName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicebus.SubscriptionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, namespaceName, topicName, subscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SubscriptionsClient) GetPreparer(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"subscriptionName":  autorest.Encode("path", subscriptionName),
		"topicName":         autorest.Encode("path", topicName),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) GetResponder(resp *http.Response) (result SBSubscription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByTopic list all the subscriptions under a specified topic.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
// topicName - the topic name.
// skip - skip is only used if a previous operation returned a partial result. If a previous response contains
// a nextLink element, the value of the nextLink element will include a skip parameter that specifies a
// starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N usageDetails.
func (client SubscriptionsClient) ListByTopic(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, skip *int32, top *int32) (result SBSubscriptionListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: topicName,
			Constraints: []validation.Constraint{{Target: "topicName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
				}}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("servicebus.SubscriptionsClient", "ListByTopic", err.Error())
	}

	result.fn = client.listByTopicNextResults
	req, err := client.ListByTopicPreparer(ctx, resourceGroupName, namespaceName, topicName, skip, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "ListByTopic", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTopicSender(req)
	if err != nil {
		result.sslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "ListByTopic", resp, "Failure sending request")
		return
	}

	result.sslr, err = client.ListByTopicResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "ListByTopic", resp, "Failure responding to request")
	}

	return
}

// ListByTopicPreparer prepares the ListByTopic request.
func (client SubscriptionsClient) ListByTopicPreparer(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, skip *int32, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"topicName":         autorest.Encode("path", topicName),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByTopicSender sends the ListByTopic request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ListByTopicSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByTopicResponder handles the response to the ListByTopic request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ListByTopicResponder(resp *http.Response) (result SBSubscriptionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByTopicNextResults retrieves the next set of results, if any.
func (client SubscriptionsClient) listByTopicNextResults(lastResults SBSubscriptionListResult) (result SBSubscriptionListResult, err error) {
	req, err := lastResults.sBSubscriptionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "listByTopicNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByTopicSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "listByTopicNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByTopicResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.SubscriptionsClient", "listByTopicNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByTopicComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubscriptionsClient) ListByTopicComplete(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, skip *int32, top *int32) (result SBSubscriptionListResultIterator, err error) {
	result.page, err = client.ListByTopic(ctx, resourceGroupName, namespaceName, topicName, skip, top)
	return
}
