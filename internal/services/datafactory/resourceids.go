package datafactory

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=DataFactory -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DataFactory/factories/facName1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=DataFlow -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DataFactory/factories/facName1/dataflows/dataflow1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=DataSet -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DataFactory/factories/facName1/datasets/dataSet1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=IntegrationRuntime -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DataFactory/factories/factory1/integrationruntimes/runtime1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=LinkedService -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DataFactory/factories/factory1/linkedservices/linkedService1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ManagedPrivateEndpoint -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DataFactory/factories/factory1/managedVirtualNetworks/vnet1/managedPrivateEndpoints/endpoint1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Trigger -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DataFactory/factories/factory1/triggers/trigger1
