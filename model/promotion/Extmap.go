package promotion

// Extmap 
type Extmap struct {
    // 扩展字段
    Keyvalue   bool `json:"keyvalue,omitempty" xml:"keyvalue,omitempty"`
    // empty
    Empty   bool `json:"empty,omitempty" xml:"empty,omitempty"`
    // 扩展字段
    Key   string `json:"key,omitempty" xml:"key,omitempty"`
}
