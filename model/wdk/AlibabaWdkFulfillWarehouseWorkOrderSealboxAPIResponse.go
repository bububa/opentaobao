package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse 仓封箱回告 API返回值
// alibaba.wdk.fulfill.warehouse.work.order.sealbox
//
// 仓封箱回告箱与包裹的关系
type AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponseModel
}

// AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponseModel is 仓封箱回告 成功返回结果
type AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_warehouse_work_order_sealbox_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败返回原因
	RespMessage string `json:"resp_message,omitempty" xml:"resp_message,omitempty"`
	// 成功失败码
	RespCode string `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
