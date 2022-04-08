package cn

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/networks/rpc"
	"io/ioutil"
	"math/big"
	"testing"
)

func loadTraceConfig() (*TraceConfig, error) {
	tracerPath := "./tracers/internal/tracers/call_tracer.js"
	tracerTimeout := "120s"
	loadedFile, err := ioutil.ReadFile(tracerPath)
	if err != nil {
		return nil, fmt.Errorf("%w: could not load tracer file", err)
	}

	loadedTracer := string(loadedFile)
	return &TraceConfig{
		Timeout: &tracerTimeout,
		Tracer:  &loadedTracer,
	}, nil
}

type rpcCall struct {
	Result *Call `json:"result"`
}

type rpcRawCall struct {
	Result json.RawMessage `json:"result"`
}

// Call is an Klaytn debug trace.
type Call struct {
	Type         string         `json:"type"`
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Value        *big.Int       `json:"value"`
	GasUsed      *big.Int       `json:"gasUsed"`
	Revert       bool
	ErrorMessage string  `json:"error"`
	Calls        []*Call `json:"calls"`
}

func TestTracerOutput(t *testing.T) {
	klaytnClient, _ := rpc.Dial("http://13.125.52.163:8551")
	tc, _ := loadTraceConfig()

	//richPrvKey := "0x275c7a80d43cab4c97b3ba7a0fb06f18321ade5de9bc0daa39d07bc03d474bc2"

	//vtBlock := "0x3576fe0fb3c32d4306adf3807cb057fa0461ab2d2155b14b6a8307652a70721e" // from(o), to(o)
	legacyBlock := "0x41906646348a64dfac2615bb8586bc7f1d35ff27f88c3fcc711de90af49d5392" // from(o), to(o)
	//deployBlock := "0xf73367c1216b9cb7c1f1d4a0a2cc7f1aa443a6d6b68d36676f53cc0721021728" // from(o), to(x)
	//exeBlock := "0x92a1c3a3e5d37e5041cc3ca7c6d3171512dbc8b625e4000394f19fddae7bfbd0" // from(o), to(o)

	var calls []*rpcCall
	var rawCalls []*rpcRawCall
	var raw json.RawMessage
	_ = klaytnClient.CallContext(context.Background(), &raw, "debug_traceBlockByHash", common.HexToHash(legacyBlock), tc)

	// This is for testing myself
	var mapCalls []*map[string]interface{}
	if err := json.Unmarshal(raw, &mapCalls); err != nil {
		panic(err)
	}

	// Decode []*rpcCall
	if err := json.Unmarshal(raw, &calls); err != nil {
		panic(err)
	}

	for _, c := range rawCalls {
		var traceMap map[string]interface{}
		_ = json.Unmarshal(c.Result, &traceMap)
		fmt.Println(traceMap)
	}
}
