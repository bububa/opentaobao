package travel

// Poi 结构体
type Poi struct {
	// POI对应的名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// POI对应ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
