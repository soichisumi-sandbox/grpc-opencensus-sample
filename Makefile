proto-build:
	protoc -I=./proto -I=./proto/ext --go_out=plugins=grpc:./app/proto --grpc-gateway_out=logtostderr=true:./app/proto proto/*.proto

go-build:
	GO111MODULE=on go build -o cmd/api/api ./cmd/api/
	GO111MODULE=on go build -o cmd/gw/gw ./cmd/gw

docker-build:
	docker build -f build/Dockerfile.api -t soichisumi0/grpc-opencensus-sample-api:v1.0 .
	docker build -f build/Dockerfile.gw -t soichisumi0/grpc-opencensus-sample-gw:v1.0 .

use-local-docker:
	eval \$\(minikube docker-env\)

minikube-start:
	minikube start

minikube-ip:
	minikube ip

prometheus-install:
	helm install --name prometheus -f deployments/helm/prometheus-customize-values.yaml stable/prometheus

prometheus-update:
	helm upgrade -f deployments/helm/prometheus-customize-values.yaml prometheus stable/prometheus

grafana-install:
	helm install --name grafana -f deployments/helm/grafana-customize-values.yaml stable/grafana

grafana-update:
	helm upgrade -f deployments/helm/grafana-customize-values.yaml grafana stable/grafana

# https://golang.org/pkg/cmd/go/internal/clean/
clean-modcache:
	go clean -modcache