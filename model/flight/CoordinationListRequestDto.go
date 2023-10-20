package flight

import (
	"sync"
)

// CoordinationListRequestDto 结构体
type CoordinationListRequestDto struct {
	// 协同单状态列表，0:“待分配”，1:“待处理”，2:“处理中”，3:“已完结”，4:“已拒绝”，5:“已关闭”，6:“待验收”，7:“预约处理”
	StatusList []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
	// 协同单类型列表，1, &#34;加订婴童&#34;； 2, &#34;加订婴童待出票&#34;； 3, &#34;自营履约系统转人工&#34;； 4, &#34;正向履约返现转人工&#34;； 5, &#34;催出票&#34;； 6, &#34;催退票&#34;； 7, &#34;催改签回填&#34;； 8, &#34;拦截出票&#34;； 9, &#34;系统主动二次回填&#34;； 10, &#34;商家主动二次回填&#34;； 13, &#34;追款&#34;； 11, &#34;改名/改证件询价&#34;； 12, &#34;改名/改证件&#34;； 4, &#34;补退&#34;； 15, &#34;宠物托运&#34;； 16, &#34;特殊餐食&#34;； 17, &#34;轮椅/担架&#34;； 18, &#34;婴儿摇篮&#34;； 19, &#34;加订行李&#34;； 20, &#34;航变清Q转人工&#34;； 22,&#34;无法出票&#34;； 23,&#34;线下编码预定&#34;； 24,&#34;申请关闭验真&#34;； 25, &#34;拦截改签&#34;； 26, &#34;催改签出票&#34;；
	CaseTypeList []string `json:"case_type_list,omitempty" xml:"case_type_list>string,omitempty"`
	// 创建时间起点
	CreateGmtBegin string `json:"create_gmt_begin,omitempty" xml:"create_gmt_begin,omitempty"`
	// 创建时间终点
	CreateGmtEnd string `json:"create_gmt_end,omitempty" xml:"create_gmt_end,omitempty"`
	// 飞猪订单号
	CorrelationOutOrderId string `json:"correlation_out_order_id,omitempty" xml:"correlation_out_order_id,omitempty"`
	// 开始页
	StartIndex int64 `json:"start_index,omitempty" xml:"start_index,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 1:国内，2:国际
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}

var poolCoordinationListRequestDto = sync.Pool{
	New: func() any {
		return new(CoordinationListRequestDto)
	},
}

// GetCoordinationListRequestDto() 从对象池中获取CoordinationListRequestDto
func GetCoordinationListRequestDto() *CoordinationListRequestDto {
	return poolCoordinationListRequestDto.Get().(*CoordinationListRequestDto)
}

// ReleaseCoordinationListRequestDto 释放CoordinationListRequestDto
func ReleaseCoordinationListRequestDto(v *CoordinationListRequestDto) {
	v.StatusList = v.StatusList[:0]
	v.CaseTypeList = v.CaseTypeList[:0]
	v.CreateGmtBegin = ""
	v.CreateGmtEnd = ""
	v.CorrelationOutOrderId = ""
	v.StartIndex = 0
	v.PageSize = 0
	v.DomesticIntl = 0
	poolCoordinationListRequestDto.Put(v)
}
