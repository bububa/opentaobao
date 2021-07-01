package travel

// PontusTravelItemItineraryInfo 结构体
type PontusTravelItemItineraryInfo struct {
	// 必填，1.使用xml标签（img和txt）进行图文混排，目前仅支持一段文字和多张图片，如果文件有多段，将会被合并 2.每段行程文字总和小于1500字
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 小于等于30字
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 必填，第一天行程为1，第二天行程为2
	ItineraryNo int64 `json:"itinerary_no,omitempty" xml:"itinerary_no,omitempty"`
}
