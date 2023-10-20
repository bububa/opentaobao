package logistic

import (
	"sync"
)

// SendResendLogisticsMsgDto 结构体
type SendResendLogisticsMsgDto struct {
	// 该运单所包含的货品列表
	GoodsItemList []GoodsItem `json:"goods_item_list,omitempty" xml:"goods_item_list>goods_item,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 发货单唯一标识
	BizId string `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
	// 物流公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 平台补发单唯一标识
	SourceId string `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 成功发货的发货单运单号，菜鸟为LP单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 运单状态（-1=废单，1=已下发）
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 主订单
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolSendResendLogisticsMsgDto = sync.Pool{
	New: func() any {
		return new(SendResendLogisticsMsgDto)
	},
}

// GetSendResendLogisticsMsgDto() 从对象池中获取SendResendLogisticsMsgDto
func GetSendResendLogisticsMsgDto() *SendResendLogisticsMsgDto {
	return poolSendResendLogisticsMsgDto.Get().(*SendResendLogisticsMsgDto)
}

// ReleaseSendResendLogisticsMsgDto 释放SendResendLogisticsMsgDto
func ReleaseSendResendLogisticsMsgDto(v *SendResendLogisticsMsgDto) {
	v.GoodsItemList = v.GoodsItemList[:0]
	v.MailNo = ""
	v.Msg = ""
	v.BizId = ""
	v.CompanyName = ""
	v.SourceId = ""
	v.OrderCode = ""
	v.Status = 0
	v.Tid = 0
	poolSendResendLogisticsMsgDto.Put(v)
}
