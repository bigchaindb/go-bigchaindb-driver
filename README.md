# go-bigchaindb-driver
Official Go BigchainDB driver

This is an implementation of a BigchainDB client written in golang. It follows the requirements set out in [BEP-13](https://github.com/bigchaindb/BEPs/tree/master/13).

## TO-DO

* Complete e2e test (WIP)
* Implement NewTransferTransaction (WIP)
* Increase test coverage
~* Add API key security option config~
* Add documentation and more detailed example package
* Benchmark tests
* Add connection pools

## Compatibility

This BigchainDB-driver is compatible with BigchainDB server v2.0 and above.

## Content

* [Installation](#installation)
   * [Example: Create a transaction](#example-create-a-transaction)
* [Documentation](#bigchaindb-documentation)
* [Authors](#authors)
* [License](#license)

## Installation

Run `go get -u github.com/bigchaindb/go-bigchaindb-driver`

## Example: Create a transaction

Please refer to github.com/bigchaindb/go-bigchaindb-driver/e2e_test.go for an example.

## BigchainDB Documentation

- [HTTP API Reference](https://docs.bigchaindb.com/projects/server/en/latest/http-client-server-api.html)
- [The Transaction Model](https://docs.bigchaindb.com/projects/server/en/latest/data-models/transaction-model.html?highlight=crypto%20conditions)
- [Inputs and Outputs](https://docs.bigchaindb.com/projects/server/en/latest/data-models/inputs-outputs.html)
- [Asset Transfer](https://docs.bigchaindb.com/projects/py-driver/en/latest/usage.html#asset-transfer)
- [All BigchainDB Documentation](https://docs.bigchaindb.com/)

## Authors

* unchain.io <dev@unchain.io>
* BigchainDB <dev@bigchaindb.com>

## License

Copyright 2018 unchain.io B.V. and BigchainDB GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
