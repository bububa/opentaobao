package lstfundbill

import (
	"sync"
)

// PagedResultDto 结构体
type PagedResultDto struct {
	// 子订单包装类
	ContentList []LstFundBillOrderDto `json:"content_list,omitempty" xml:"content_list>lst_fund_bill_order_dto,omitempty"`
	// 错误信息，success为false时生效
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误代码，success为false时生效
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 总记录数，请求参数page大于实际值时，total返回0
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 每页记录数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPagedResultDto = sync.Pool{
	New: func() any {
		return new(PagedResultDto)
	},
}

// GetPagedResultDto() 从对象池中获取PagedResultDto
func GetPagedResultDto() *PagedResultDto {
	return poolPagedResultDto.Get().(*PagedResultDto)
}

// ReleasePagedResultDto 释放PagedResultDto
func ReleasePagedResultDto(v *PagedResultDto) {
	v.ContentList = v.ContentList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Total = 0
	v.Size = 0
	v.Page = 0
	v.Success = false
	poolPagedResultDto.Put(v)
}
