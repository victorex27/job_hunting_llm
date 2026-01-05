package llm

type Options struct {
	Temperature *float64 `json:"temperature,omitempty"`
}

type GenerateRequest struct {
	Model   string   `json:"model"`
	Prompt  string   `json:"prompt"`
	Stream  bool     `json:"stream"` // We'll set this to false for now to keep it simple
	Options *Options `json:"options,omitempty"` // pointer allows nil/omission
	Format  string   `json:"format"`
}



type GenerateResponse struct {
	Response string `json:"response"`
	Duration int64 `json:"duration"`
	Done    bool   `json:"done"`	
}
