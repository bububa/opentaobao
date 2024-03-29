package alitripmerchant

import (
	"sync"
)

// PageableResponse 结构体
type PageableResponse struct {
	// 返回类型
	Contents []HotelListSearchDto `json:"contents,omitempty" xml:"contents>hotel_list_search_dto,omitempty"`
	// 查询抽奖用户数据集合
	Content []LotteryDataVo `json:"content,omitempty" xml:"content>lottery_data_vo,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 当前页的数量
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 当前页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页的数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	TotalPageNum int64 `json:"total_page_num,omitempty" xml:"total_page_num,omitempty"`
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否有下一页
	HasNextPage bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPageableResponse = sync.Pool{
	New: func() any {
		return new(PageableResponse)
	},
}

// GetPageableResponse() 从对象池中获取PageableResponse
func GetPageableResponse() *PageableResponse {
	return poolPageableResponse.Get().(*PageableResponse)
}

// ReleasePageableResponse 释放PageableResponse
func ReleasePageableResponse(v *PageableResponse) {
	v.Contents = v.Contents[:0]
	v.Content = v.Content[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Size = 0
	v.PageNo = 0
	v.PageSize = 0
	v.TotalPageNum = 0
	v.TotalCount = 0
	v.HasNextPage = false
	v.Success = false
	poolPageableResponse.Put(v)
}
