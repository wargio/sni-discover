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
$ ./sni-discover -target 1.1.1.1 -best 5 -h2only -debug
2024/05/06 19:04:45 Resolving 1000 SNIs.
2024/05/06 19:04:45 Waiting for resolvers.
2024/05/06 19:04:56 Sorting 1000 SNIs.
+---+---------------------+---------+----------+--------------+-------+
| # | SNI                 | TLS     | PROTO    |     DURATION | EXTRA |
+---+---------------------+---------+----------+--------------+-------+
| 0 | 020806.xyz          | TLSv1.3 | HTTP/2.0 | 486.304852ms |       |
| 1 | 111984.xyz          | TLSv1.3 | HTTP/2.0 | 622.392936ms |       |
| 2 | 111331.xyz          | TLSv1.3 | HTTP/2.0 | 654.457576ms |       |
| 3 | 0061681.xyz         | TLSv1.3 | HTTP/2.0 | 765.647363ms |       |
| 4 | 101homebusiness.com | TLSv1.3 | HTTP/2.0 | 838.300103ms |       |
+---+---------------------+---------+----------+--------------+-------+
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