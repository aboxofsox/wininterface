# WinInterface
A compatibility layer to get Windows network interface information that is compatible with gopacket/pcap.

```go
cmd := wininterface.GetMac()
macs := cmd.Parse()

for _, m := range macs {
	// handle information
}
```
