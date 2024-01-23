package lib

type Formula struct {
	Commands []Commands `json:"commands"`
}

type Commands struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Action      string `json:"action"`
}
