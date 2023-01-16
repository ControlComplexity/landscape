#protoc --go_out=. .\api\landscape\v1\landscape.proto
#protoc --go-grpc_out=. .\api\landscape\v1\landscape.proto
#protoc --grpc-gateway_opt=logtostderr=true --grpc-gateway_out=. .\api\landscape\v1\landscape.proto
#protoc --openapiv2_opt=logtostderr=true,json_names_for_fields=false --openapiv2_out=./ .\api\landscape\v1\landscape.proto
#protoc  --validate_out="lang=go:./"  .\api\landscape\v1\landscape.proto


# 生成 proto 模板
kratos proto add api/landscape/v1/exhibition.proto
# 生成 client 源码
kratos proto client api/landscape/v1/exhibition.proto
# 生成 server 源码
kratos proto server api/landscape/v1/exhibition.proto -t internal/service
#运行项目
kratos run
