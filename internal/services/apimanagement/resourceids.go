package apimanagement

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Api -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/apis/api1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ApiDiagnostic -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/apis/api1/diagnostics/diagnostic1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ApiManagement -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ApiOperation -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/apis/api1/operations/operation1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ApiOperationPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/apis/api1/operations/operation1/policies/policy1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ApiPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/apis/api1/policies/policy1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ApiSchema -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/apis/api1/schemas/schema1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ApiVersionSet -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/apiVersionSets/apiVersionSet1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=AuthorizationServer -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/authorizationServers/authorizationserver1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Backend -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/backends/backend1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=RedisCache -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/caches/redisCache1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Certificate -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/certificates/certificate1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=CustomDomain -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/customDomains/customdomain
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Diagnostic -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/diagnostics/diagnostic1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=EmailTemplate -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/templates/template1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Gateway -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=GatewayApi -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1/apis/api1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Group -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/groups/group1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=GroupUser -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/groups/group1/users/user1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=IdentityProvider -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/identityProviders/identityProvider1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Logger -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/loggers/logger1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=NamedValue -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/namedValues/namedValue1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=OpenIDConnectProvider -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/openidConnectProviders/opid1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Policy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/policies/policy1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Product -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/products/product1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ProductApi -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/products/product1/apis/api1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ProductGroup -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/products/product1/groups/group1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ProductPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/products/product1/policies/policy1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Property -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/namedValues/namedvalue1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Subscription -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/subscriptions/subscription1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=User -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/users/user1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=OperationTag -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/apis/api1/operations/operation1/tags/tag1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ApiRelease -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/apis/api1/releases/release1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Tag -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/tags/tag1
