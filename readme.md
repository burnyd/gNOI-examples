# gNOI examples in go.

Here are some examples of executing [gNOI](https://github.com/openconfig/gnoi) services.

I have a few very simple examples.

How to run each one will have something commeneted out.  For example within the system-services/ping/main.go

```text
//go run main.go -username admin -password admin -target 172.20.20.2:6030 -destination 1.1.1.1
```

#### verify os-services/verify
Verify will verify the version of code on the device.

### ping system-services
This will ping a destination

### reboot system-services
Do not use this unless you want to reboot!  Cuation!  This will reboot a device.

### traceroute
This will traceroute to a given destination.