package xiami

// BillboardItemVO 结构体
type BillboardItemVO struct {
	// songs
	Songs []Songs `json:"songs,omitempty" xml:"songs>songs,omitempty"`
	// billboardId
	BillboardId int64 `json:"billboard_id,omitempty" xml:"billboard_id,omitempty"`
}
