export PATH="$PATH:$(go env GOPATH)/bin"

### SYSTEM MODULE ###
# vault
echo "Generating proto for system_vault service"
mkdir -p ./modules/system/proto/vault
protoc --go_out=./modules/system/proto/vault --go_opt=paths=source_relative --go-grpc_out=./modules/system/proto/vault --go-grpc_opt=paths=source_relative -I ./openbp/modules/system/proto vault.proto

### NATIVE MODULE ###

# namespace
echo "Generating proto for native_namespace service"
mkdir -p ./modules/native/proto/namespace
protoc --go_out=./modules/native/proto/namespace --go_opt=paths=source_relative --go-grpc_out=./modules/native/proto/namespace --go-grpc_opt=paths=source_relative -I ./openbp/modules/native/proto namespace.proto
# keyvaluestorage
echo "Generating proto for native_keyvaluestorage service"
mkdir -p ./modules/native/proto/keyvaluestorage
protoc --go_out=./modules/native/proto/keyvaluestorage --go_opt=paths=source_relative --go-grpc_out=./modules/native/proto/keyvaluestorage --go-grpc_opt=paths=source_relative -I ./openbp/modules/native/proto keyvaluestorage.proto

# actor_user
echo "Generating proto for native_actor_user service"
mkdir -p ./modules/native/proto/actor/user
protoc --go_out=./modules/native/proto/actor/user --go_opt=paths=source_relative --go-grpc_out=./modules/native/proto/actor/user --go-grpc_opt=paths=source_relative -I ./openbp/modules/native/proto/actor user.proto

# iam_token
echo "Generating proto for native_iam_token service"
mkdir -p ./modules/native/proto/iam/token
protoc --go_out=./modules/native/proto/iam/token --go_opt=paths=source_relative --go-grpc_out=./modules/native/proto/iam/token --go-grpc_opt=paths=source_relative -I ./openbp/modules/native/proto/iam token.proto
# iam_policy
echo "Generating proto for native_iam_policy service"
mkdir -p ./modules/native/proto/iam/policy
protoc --go_out=./modules/native/proto/iam/policy --go_opt=paths=source_relative --go-grpc_out=./modules/native/proto/iam/policy --go-grpc_opt=paths=source_relative -I ./openbp/modules/native/proto/iam policy.proto
# iam_role
echo "Generating proto for native_iam_role service"
mkdir -p ./modules/native/proto/iam/role
protoc --go_out=./modules/native/proto/iam/role --go_opt=paths=source_relative --go-grpc_out=./modules/native/proto/iam/role --go-grpc_opt=paths=source_relative -I ./openbp/modules/native/proto/iam role.proto
# iam_auth
echo "Generating proto for native_iam_auth service"
mkdir -p ./modules/native/proto/iam/auth
protoc --go_out=./modules/native/proto/iam/auth --go_opt=paths=source_relative --go-grpc_out=./modules/native/proto/iam/auth --go-grpc_opt=paths=source_relative -I ./openbp/modules/native/proto/iam auth.proto
# iam_identity
echo "Generating proto for native_iam_identity service"
mkdir -p ./modules/native/proto/iam/identity
protoc --go_out=./modules/native/proto/iam/identity --go_opt=paths=source_relative --go-grpc_out=./modules/native/proto/iam/identity --go-grpc_opt=paths=source_relative -I ./openbp/modules/native/proto/iam identity.proto
# iam_authentication_password
echo "Generating proto for native_iam_authentication_password service"
mkdir -p ./modules/native/proto/iam/authentication/password
protoc --go_out=./modules/native/proto/iam/authentication/password --go_opt=paths=source_relative --go-grpc_out=./modules/native/proto/iam/authentication/password --go-grpc_opt=paths=source_relative -I ./openbp/modules/native/proto/iam/authentication password.proto