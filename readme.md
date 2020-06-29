## Tomato Exporter

Tomato exporter is a simple app that
scrapes advanced tomato routers and 
dumps the bandwidth usage into prometheus.

### Config

To see an example config, see the included
`example.yaml`, which documents all the available
options.

### Building

This project uses goxz to bundle and cross compile.

```bash
goxz -d dist -include example.conf,example.service cmd/tomato_exporter/main.go
```
