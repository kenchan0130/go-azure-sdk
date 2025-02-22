
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/dataversion` Documentation

The `dataversion` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/dataversion"
```


### Client Initialization

```go
client := dataversion.NewDataVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataVersionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := dataversion.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "dataValue", "versionValue")

payload := dataversion.DataVersionBaseResource{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataVersionClient.Delete`

```go
ctx := context.TODO()
id := dataversion.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "dataValue", "versionValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataVersionClient.Get`

```go
ctx := context.TODO()
id := dataversion.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "dataValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataVersionClient.List`

```go
ctx := context.TODO()
id := dataversion.NewDataID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "dataValue")

// alternatively `client.List(ctx, id, dataversion.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, dataversion.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
