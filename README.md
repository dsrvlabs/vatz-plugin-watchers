# vatz-plugin-watcher-cosmos
Vatz plugin for monitoring node's validating status 


## Plugins
- watcher_cosmos : monitor validating signatures in blocks

## Installation and Usage
> Please make sure [Vatz](https://github.com/dsrvlabs/vatz) is running with proper configuration. [Vatz Installation Guide](https://github.com/dsrvlabs/vatz/blob/main/docs/installation.md)

### Install Plugins
- Install with source
```
$ git clone https://github.com/dsrvlabs/vatz-plugin-watcher-cosmos.git
$ cd vatz-plugin-watcher-cosmos
$ make install
```

- Install with Vatz CLI command
```
$ vatz plugin install --help
Install new plugin

Usage:
   plugin install [flags]

Examples:
vatz plugin install github.com/dsrvlabs/<somewhere> name

Flags:
  -h, --help   help for install
```
> please make sure install path for the plugins repository URL.
```
$ vatz plugin install github.com/dsrvlabs/vatz-plugin-watcher-cosmos/plugins/watcher_cosmos node_watcher_cosmos
```
- Check plugins list with Vatz CLI command
```
$ vatz plugin list                                                                                                                                            
+----------------+------------+---------------------+-------------------------------------------------------------------------------+---------+
| NAME           | IS ENABLED | INSTALL DATE        | REPOSITORY                                                                    | VERSION |
+----------------+------------+---------------------+-------------------------------------------------------------------------------+---------+
| watcher_cosmos | true       | 2024-07-19 12:26:50 | github.com/dsrvlabs/vatz-plugin-watcher-cosmos/plugins/watcher_cosmos         | latest  |
+----------------+------------+---------------------+-------------------------------------------------------------------------------+---------+

```

### Run
> Run as default config or option flags
```
$ watcher_cosmos
2024-07-19T12:28:40-05:00 INF Register module=grpc
2024-07-19T12:28:40-05:00 INF Start 127.0.0.1 10001 module=sdk
2024-07-19T12:28:40-05:00 INF Start module=grpc
2024-07-19T12:29:09-05:00 INF Execute module=grpc
2024-07-19T12:29:09-05:00 DBG The validator is signing the block successfully. module=plugin
2024-07-19T12:29:18-05:00 INF Execute module=grpc
```


## Command line arguments
- node_block_sync
```
Usage of node_block_sync:
  -addr string
	Listening address (default "127.0.0.1")
  -port int
	Listening port (default 10001)
  -rpcURI string
	Tendermint RPC URI Address (default "http://localhost:26657")
  -voterAddr(Hex) string
    Need to Validator Operator Address (Hex) (mendatory)
  -warning int
    block height stucked count to raise warning level of alert (default 3)
  -critical int
	block height stucked count to raise critical level of alert (default 10)
```

## Using Script for plugins mandatory flags
```
$ ./sciprt/get_veloper_addre_hex.sh
Enter the Rest Endpoint: http://127.0.0.1:26657
Enter the RPC Endpoint: http://127.0.0.1:26657
Enter the valoper address: someValoperAddr1tt8dczjk62pnwwanm99rq2kw25
 
You have entered the following details:
Rest Endpoint: http://127.0.0.1:26657
RPC Endpoint: http://127.0.0.1:26657
Valoper(Validator Operator) address: someValoperAddr1tt8dczjk62pnwwanm99rq2kw25
Are the entered values correct? Do you wish to proceed? (y/n): 
Validator address: cosmosvaloper1...
```

## TroubleShooting
1. Encountered issue related with `Device or Resource Busy` or `Too many open files` error.
- Check your open file limit and recommended to increase it.
 ```
 $ ulimit -n
 1000000
 ```

## License

`vatz-plugin-watcher-cosmos` is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included in our repository in the `LICENSE` file.
