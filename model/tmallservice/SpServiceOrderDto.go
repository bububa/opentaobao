package tmallservice

import (
	"sync"
)

// SpServiceOrderDto 结构体
type SpServiceOrderDto struct {
	// 费用信息
	FeeList []FeeInfo `json:"fee_list,omitempty" xml:"fee_list>fee_info,omitempty"`
	// 服务过期时间
	GmtExpire string `json:"gmt_expire,omitempty" xml:"gmt_expire,omitempty"`
	// 服务单修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 服务单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 服务单有效期开始时间
	GmtEffect string `json:"gmt_effect,omitempty" xml:"gmt_effect,omitempty"`
	// 状态编码：create(创建)、effect(生效)、closed(关闭)、finish(完成)
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 履约类型:1, &#34;到店&#34;2, &#34;到家&#34;3, &#34;寄送&#34;
	FulfilTypeCode string `json:"fulfil_type_code,omitempty" xml:"fulfil_type_code,omitempty"`
	// 取消的份数
	RefundServiceCount int64 `json:"refund_service_count,omitempty" xml:"refund_service_count,omitempty"`
	// 实物子订单信息
	MasterTradeOrder *MasterTradeOrderDto `json:"master_trade_order,omitempty" xml:"master_trade_order,omitempty"`
	// 服务定义
	ServiceDefinition *ServiceDefinitionDto `json:"service_definition,omitempty" xml:"service_definition,omitempty"`
	// 买家信息
	Buyer *BuyerDto `json:"buyer,omitempty" xml:"buyer,omitempty"`
	// 已使用份数
	UsedServiceCount int64 `json:"used_service_count,omitempty" xml:"used_service_count,omitempty"`
	// 剩余的份数
	LeftServiceCount int64 `json:"left_service_count,omitempty" xml:"left_service_count,omitempty"`
	// 服务子订单信息
	ServiceTradeOrder *ServiceTradeOrderDto `json:"service_trade_order,omitempty" xml:"service_trade_order,omitempty"`
	// 服务的总份数
	ServiceCount int64 `json:"service_count,omitempty" xml:"service_count,omitempty"`
	// 服务单id
	SpServiceorderId int64 `json:"sp_serviceorder_id,omitempty" xml:"sp_serviceorder_id,omitempty"`
	// 正在使用中的次数
	UsingServiceCount int64 `json:"using_service_count,omitempty" xml:"using_service_count,omitempty"`
	// 服务定义
	ServiceDefinitionDTO *ServiceDefinitionDto `json:"service_definition_d_t_o,omitempty" xml:"service_definition_d_t_o,omitempty"`
	// 服务单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 服务交易订单
	ServiceTradeOrderDTO *ServiceTradeOrderDto `json:"service_trade_order_d_t_o,omitempty" xml:"service_trade_order_d_t_o,omitempty"`
	// 服务单申请工单的幂等键
	ServiceSequence int64 `json:"service_sequence,omitempty" xml:"service_sequence,omitempty"`
	// 服务提供者
	ServiceProviderDTO *ServiceProviderDto `json:"service_provider_d_t_o,omitempty" xml:"service_provider_d_t_o,omitempty"`
}

var poolSpServiceOrderDto = sync.Pool{
	New: func() any {
		return new(SpServiceOrderDto)
	},
}

// GetSpServiceOrderDto() 从对象池中获取SpServiceOrderDto
func GetSpServiceOrderDto() *SpServiceOrderDto {
	return poolSpServiceOrderDto.Get().(*SpServiceOrderDto)
}

// ReleaseSpServiceOrderDto 释放SpServiceOrderDto
func ReleaseSpServiceOrderDto(v *SpServiceOrderDto) {
	v.FeeList = v.FeeList[:0]
	v.GmtExpire = ""
	v.GmtModified = ""
	v.GmtCreate = ""
	v.GmtEffect = ""
	v.StatusCode = ""
	v.FulfilTypeCode = ""
	v.RefundServiceCount = 0
	v.MasterTradeOrder = nil
	v.ServiceDefinition = nil
	v.Buyer = nil
	v.UsedServiceCount = 0
	v.LeftServiceCount = 0
	v.ServiceTradeOrder = nil
	v.ServiceCount = 0
	v.SpServiceorderId = 0
	v.UsingServiceCount = 0
	v.ServiceDefinitionDTO = nil
	v.Id = 0
	v.ServiceTradeOrderDTO = nil
	v.ServiceSequence = 0
	v.ServiceProviderDTO = nil
	poolSpServiceOrderDto.Put(v)
}
