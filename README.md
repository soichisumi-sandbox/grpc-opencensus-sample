# grpc-opencensus-sample

sample implementation of grpc server with metrics collection & distributed tracing by [OpenCensus](https://opencensus.io/)

## Usage

### Build

* `make proto-build`
* `make go-build`

### Run

* execute `cmd/api/api`

### Request example

1. launch [Evans](https://github.com/ktr0731/evans) 
    * evans -r --host localhost -p 3000
2. select service
    * service Server
3. call rpc
    * call GetData