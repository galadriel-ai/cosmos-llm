# cosmosllm
**cosmosllm** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).


## Useful commands

```
# Make inference request
cosmos-llmd tx inference run-inference "hello3" 1 --from bob --chain-id cosmosllm --fees 5000stake
# Get inference request
cosmos-llmd q inference get-inference-run 0 --chain-id cosmosllm

# Join "gpu" pool
cosmos-llmd tx inference join-inference-pool 1 --from bob  --chain-id cosmosllm --fees 5000stake
# Get pool size
cosmos-llmd q inference get-pool-size --chain-id cosmosllm

# Get unanswered query by model ID
cosmos-llmd q inference get-unanswered-inference-run 1 --chain-id cosmosllm
```

Get acc balance
```
curl \
    -X GET \
    -H "Content-Type: application/json" \
    http://localhost:1317/cosmos/bank/v1beta1/balances/cosmos1xt2c7xnce4p3npvpewvkt4j50mnzu3lz8duuas
```

**scaffolding commands**
```
# Write endpoint
ignite scaffold message runInference prompt modelid:uint --response id:uint --module inference

# query endpoint
ignite scaffold query get-inference-run id:uint --response inference:Inferencerun --module inference

# Data type
ignite scaffold type inferencerun modelId:uint prompt:string isfinished:bool id:uint --module inference

# Regenerate boilerplate for some type after changing .proto file(s)
ignite generate proto-go
```

## Testing
```
go install go.uber.org/mock/mockgen@latest
```

**Generate mocks**
```
mockgen -source=x/inference/types/expected_keepers.go \
                -package testutil \
                -destination=x/inference/testutil/expected_keepers_mocks.go
```
**Run some tests**
```
go test -v ./x/inference/keeper
```

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/cosmos-llm@latest! | sudo bash
```
`username/cosmos-llm` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
