package omniorder

// OrderDetailDto 结构体
type OrderDetailDto struct {
	// 取件信息
	PickUpInfos []Content `json:"pick_up_infos,omitempty" xml:"pick_up_infos>content,omitempty"`
}
