package maitix

// IdNameDTO 
type IdNameDTO struct {
    // 城市id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 城市名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
