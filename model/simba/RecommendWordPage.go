package simba

import (
	"sync"
)

// RecommendWordPage 结构体
type RecommendWordPage struct {
	// 推荐词分页对象列表
	RecommendWordList []RecommendWord `json:"recommend_word_list,omitempty" xml:"recommend_word_list>recommend_word,omitempty"`
	// 每页数据大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 返回第几页的数据从1开始。&lt;br/&gt;如果输入的值大于可取得的最大页码值时，将返回&lt;br/&gt;最大的页码值并且recommend_word_list值将&lt;br/&gt;为空
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 所查询的数据总数
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
}

var poolRecommendWordPage = sync.Pool{
	New: func() any {
		return new(RecommendWordPage)
	},
}

// GetRecommendWordPage() 从对象池中获取RecommendWordPage
func GetRecommendWordPage() *RecommendWordPage {
	return poolRecommendWordPage.Get().(*RecommendWordPage)
}

// ReleaseRecommendWordPage 释放RecommendWordPage
func ReleaseRecommendWordPage(v *RecommendWordPage) {
	v.RecommendWordList = v.RecommendWordList[:0]
	v.PageSize = 0
	v.PageNo = 0
	v.TotalItem = 0
	poolRecommendWordPage.Put(v)
}
