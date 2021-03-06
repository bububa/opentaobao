package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailCommissionOrderSyncAPIResponse 分佣数据传输 API返回值
// alibaba.retail.commission.order.sync
//
// 同步分佣结果
type AlibabaRetailCommissionOrderSyncAPIResponse struct {
	model.CommonResponse
	AlibabaRetailCommissionOrderSyncAPIResponseModel
}

// AlibabaRetailCommissionOrderSyncAPIResponseModel is 分佣数据传输 成功返回结果
type AlibabaRetailCommissionOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_commission_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的执行状态吗
	SCode string `json:"s_code,omitempty" xml:"s_code,omitempty"`
	// 返回的数据实体
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否执行成功
	SSuccess bool `json:"s_success,omitempty" xml:"s_success,omitempty"`
}
