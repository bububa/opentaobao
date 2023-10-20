package wdk

import (
	"sync"
)

// RefundAggregateQueryResult 结构体
type RefundAggregateQueryResult struct {
	// 退款单id列表
	RefundIdList []int64 `json:"refund_id_list,omitempty" xml:"refund_id_list>int64,omitempty"`
	// 淘鲜达子订单id列表
	BizIdList []int64 `json:"biz_id_list,omitempty" xml:"biz_id_list>int64,omitempty"`
	// 淘宝子订单id列表
	TbBizIdList []int64 `json:"tb_biz_id_list,omitempty" xml:"tb_biz_id_list>int64,omitempty"`
	// 接口返回码. 如果返回 HM05038888888006 需重试(数据查询失败，请重试，注意限定重试次数)
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 接口返回码描述
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 结果数量
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 下一页序号
	NextIndex int64 `json:"next_index,omitempty" xml:"next_index,omitempty"`
}

var poolRefundAggregateQueryResult = sync.Pool{
	New: func() any {
		return new(RefundAggregateQueryResult)
	},
}

// GetRefundAggregateQueryResult() 从对象池中获取RefundAggregateQueryResult
func GetRefundAggregateQueryResult() *RefundAggregateQueryResult {
	return poolRefundAggregateQueryResult.Get().(*RefundAggregateQueryResult)
}

// ReleaseRefundAggregateQueryResult 释放RefundAggregateQueryResult
func ReleaseRefundAggregateQueryResult(v *RefundAggregateQueryResult) {
	v.RefundIdList = v.RefundIdList[:0]
	v.BizIdList = v.BizIdList[:0]
	v.TbBizIdList = v.TbBizIdList[:0]
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.TotalNum = 0
	v.NextIndex = 0
	poolRefundAggregateQueryResult.Put(v)
}
