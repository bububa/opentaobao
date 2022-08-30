package wlbimports

// BigbagCreateResponse 结构体
type BigbagCreateResponse struct {
	// 大包Code
	BigbagCode string `json:"bigbag_code,omitempty" xml:"bigbag_code,omitempty"`
	// 大包id
	BigbagId int64 `json:"bigbag_id,omitempty" xml:"bigbag_id,omitempty"`
}
