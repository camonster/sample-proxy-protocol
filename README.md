# sample-proxy-protocol

- 通常のnet.Listenerを利用した場合

```bash
$ curl 10.194.229.195
RemoteAddr: 10.194.229.254:15059
```

- proxyproto.Listenerでラップしたものを利用した場合

```bash
$ curl 10.194.229.195
RemoteAddr: 10.48.2.4:54154
```
