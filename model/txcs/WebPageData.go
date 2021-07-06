package txcs

// WebPageData 结构体
type WebPageData struct {
	// 结果数据
	List []WebPageDataList `json:"list,omitempty" xml:"list>web_page_data_list,omitempty"`
	// 页码信息
	Pagination *Pagination `json:"pagination,omitempty" xml:"pagination,omitempty"`
}
