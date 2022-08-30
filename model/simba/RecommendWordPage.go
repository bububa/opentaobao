package simba

// RecommendWordPage 结构体
type RecommendWordPage struct {
	// 推荐词分页对象列表
	RecommendWordList []RecommendWord `json:"recommend_word_list,omitempty" xml:"recommend_word_list>recommend_word,omitempty"`
	// 每页数据大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 返回第几页的数据从1开始。<br/>如果输入的值大于可取得的最大页码值时，将返回<br/>最大的页码值并且recommend_word_list值将<br/>为空
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 所查询的数据总数
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
}
