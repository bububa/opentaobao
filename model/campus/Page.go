package campus

import (
	"sync"
)

// Page 结构体
type Page struct {
	// 返回内容
	Items []UserRolesDto `json:"items,omitempty" xml:"items>user_roles_dto,omitempty"`
	// 分页结果集合
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// 分页集合
	ResultList []DeviceStandardApiDto `json:"result_list,omitempty" xml:"result_list>device_standard_api_dto,omitempty"`
	// 结果
	Result []string `json:"result,omitempty" xml:"result>string,omitempty"`
	// 总计数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 每页多少条
	Limit int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolPage = sync.Pool{
	New: func() any {
		return new(Page)
	},
}

// GetPage() 从对象池中获取Page
func GetPage() *Page {
	return poolPage.Get().(*Page)
}

// ReleasePage 释放Page
func ReleasePage(v *Page) {
	v.Items = v.Items[:0]
	v.Results = v.Results[:0]
	v.ResultList = v.ResultList[:0]
	v.Result = v.Result[:0]
	v.Total = 0
	v.PageSize = 0
	v.CurrentPage = 0
	v.TotalPage = 0
	v.Limit = 0
	v.TotalCount = 0
	poolPage.Put(v)
}
