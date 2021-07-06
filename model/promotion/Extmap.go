package promotion

// Extmap 结构体
type Extmap struct {
	// 扩展字段
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 扩展字段
	Keyvalue bool `json:"keyvalue,omitempty" xml:"keyvalue,omitempty"`
	// empty
	Empty bool `json:"empty,omitempty" xml:"empty,omitempty"`
}
