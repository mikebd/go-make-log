package main

import "encoding/json"

func jsonLogLine() string {
	// TODO: Monotonic increasing timestamp
	// TODO: Randomize level
	// TODO: Randomize message
	result, _ := json.Marshal(map[string]interface{}{
		"timestamp": "2021-10-13T18:02:00.000Z",
		"level":     "info",
		"message":   "GET /api/v1/applyboard/programs/1234",
	})

	return string(result)
}
