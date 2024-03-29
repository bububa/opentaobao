package ascpchannel

import (
	"sync"
)

// AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData 结构体
type AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData struct {
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 期货库存账户id
	AicChannelInvId string `json:"aic_channel_inv_id,omitempty" xml:"aic_channel_inv_id,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 外部交易号(主)
	ExtOrderId string `json:"ext_order_id,omitempty" xml:"ext_order_id,omitempty"`
	// 外部交易号(子)
	ExtSubOrderId string `json:"ext_sub_order_id,omitempty" xml:"ext_sub_order_id,omitempty"`
	// 物流货主ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 调用接口是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData = sync.Pool{
	New: func() any {
		return new(AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData)
	},
}

// GetAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData() 从对象池中获取AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData
func GetAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData() *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData {
	return poolAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData.Get().(*AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData)
}

// ReleaseAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData 释放AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData
func ReleaseAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData(v *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData) {
	v.StoreCode = ""
	v.ChannelCode = ""
	v.AicChannelInvId = ""
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.ExtOrderId = ""
	v.ExtSubOrderId = ""
	v.UserId = 0
	v.ScItemId = 0
	v.Success = false
	poolAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateData.Put(v)
}
