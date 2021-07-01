package xiami

// BillboardItemVO 结构体
type BillboardItemVO struct {
	// billboardId
	BillboardId int64 `json:"billboard_id,omitempty" xml:"billboard_id,omitempty"`
	// songs
	Songs []Songs `json:"songs,omitempty" xml:"songs>songs,omitempty"`
}
