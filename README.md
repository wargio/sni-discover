# SNI Discovery

Quickly discover multiple Server Name Indication (SNI) within the same Autonomous System (AS).

The domains selected by this tool meet the following criteria:

* They are under the same AS as the target IP address, ensuring IP similarity.
* They support TLS version 1.3.
* They support HTTP/2 protocol (optional).

## How to Use

To use this tool, follow these steps:

1. Go to the release page and download the executable file.
2. Execute the tool as follows

```
./sni-discover -target 1.2.3.4 -best 10 -h2only
```

Optional parameters include:
```
Usage of ./sni-discover:
  -best int
        Shows the best N results only when non-zero.
  -debug
        When enabled, shows additional information.
  -file string
        File where to write the SNIs.
  -h2only
        When enabled only shows HTTP2 only results.
  -routines int
        Maximum number of routines. (default 64)
  -target string
        Target IPv4 address to use for discovering SNIs.
  -timeout int
        Connection timeout in seconds. (default 10)

```

## More

This tool uses [BGP.he.net](https://bgp.he.net) for DNS discovery

## License

This tool is licensed under the MIT license.