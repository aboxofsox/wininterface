package wininterface

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Cmd string

// WinInterface represents the output from getmac
type WinInterface struct {
	ConnectionName  string
	NetworkAdapter  string
	PhysicalAddress string
	TransportName   string
}

// Because Windows
const (
	CR = "\r" // carriage return
	LF = "\n" // line feed
)

// GetMac runs the getmac Windows command and returns its output.
func GetMac() Cmd {
	if _, err := exec.LookPath("getmac"); err != nil {
		fmt.Println("getmac is not installed")
		os.Exit(1)
	}

	cmd := exec.Command("getmac", "/FO", "list", "/V")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	return Cmd(output)
}

// chunkSlice splits a slice by "chunk" given a chunk size.
func chunkSlice(slice []string, size int) [][]string {
	var chunks [][]string

	for i := 0; i < len(slice); i += size {
		end := i + size

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

// chunkMap splits a map by "chunk" given a chunk size
func chunkMap(mp []map[string]string, size int) [][]map[string]string {
	var chunks [][]map[string]string
	for i := 0; i < len(mp); i += size {
		end := i + size
		if end > len(mp) {
			end = len(mp)
		}

		chunks = append(chunks, mp[i:end])
	}

	return chunks
}

// Parse takes the output of GetMac as a list and parses it for use, because Windows
func (cmd Cmd) Parse() []WinInterface {
	var amp []map[string]string
	list := string(cmd)

	listSplit := strings.Split(list, CR+LF)
	listChunks := chunkSlice(listSplit, 5)

	for i := range listChunks {
		mp := make(map[string]string)
		for j := range listChunks[i] {
			v := listChunks[i][j]

			vs := strings.Split(v, ":")
			if len(vs) == 2 {
				mp[vs[0]] = strings.Trim(vs[1], " ")
			}
		}
		amp = append(amp, mp)
	}

	dataChunks := chunkMap(amp, 4)

	var wis []WinInterface

	for i := range dataChunks {
		var wi WinInterface

		for _, chunk := range dataChunks[i] {

			if chunk["Connection Name"] != "" {
				wi.ConnectionName = chunk["Connection Name"]
				wi.NetworkAdapter = chunk["Network Adapter"]
				wi.PhysicalAddress = chunk["Physical Address"]

				tp := strings.ReplaceAll(chunk["Transport Name"], "Tcpip", "NPF")

				wi.TransportName = strings.ToLower(tp)

				wis = append(wis, wi)
			}

		}
	}

	return wis

}
