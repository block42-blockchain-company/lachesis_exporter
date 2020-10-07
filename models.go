package main

type ResponseBody struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int64  `json:"id"`
	Result  string `json:"result"`
}
type ResponseBodyHead struct {
	JSONRPC string   `json:"jsonrpc"`
	ID      int64    `json:"id"`
	Result  []string `json:"result"`
}
type ResponseBodyHeadInterface struct {
	JSONRPC string                 `json:"jsonrpc"`
	ID      int64                  `json:"id"`
	Result  map[string]interface{} `json:"result"`
}

type RequestBody struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      int64  `json:"id"`
}

// Probably not necessary....
type IntParamRequestBody struct {
	JSONRPC string  `json:"jsonrpc"`
	Method  string  `json:"method"`
	ID      int64   `json:"id"`
	Params  []int64 `json:"params"`
}

type StringParamRequestBody struct {
	JSONRPC string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	ID      int64    `json:"id"`
	Params  []string `json:"params"`
}

type InterfaceParamRequestBody struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	ID      int64         `json:"id"`
	Params  []interface{} `json:"params"`
}

type DownTimeResponse struct {
	JSONRPC string         `json:"jsonrpc"`
	ID      int64          `json:"id"`
	Result  DownTimeResult `json:"result"`
}

type DownTimeResult struct {
	Downtime     string `json:"downtime"`
	MissedBlocks string `json:"missedBlocks"`
}
