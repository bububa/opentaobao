package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslogisticscreatewarehouseorderAPIResponse 创建线上物流订单 API返回值
// aliexpress.logistics.createwarehouseorder
//
// 创建线上发货物流订单
type AliexpresslogisticscreatewarehouseorderAPIResponse struct {
	model.CommonResponse
	AliexpresslogisticscreatewarehouseorderAPIResponseModel
}

// AliexpresslogisticscreatewarehouseorderAPIResponseModel is 创建线上物流订单 成功返回结果
type AliexpresslogisticscreatewarehouseorderAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_logistics_createwarehouseorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AeopWlCreateWarehouseOrderResultDto `json:"result,omitempty" xml:"result,omitempty"`
	// 调用是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
