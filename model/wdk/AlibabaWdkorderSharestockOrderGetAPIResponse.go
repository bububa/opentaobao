package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkordersharestockordergetAPIResponse 猫超商户订单拉取 API返回值
// alibaba.wdkorder.sharestock.order.get
//
// 商户拉取猫超订单数据
type AlibabawdkordersharestockordergetAPIResponse struct {
	model.CommonResponse
	AlibabawdkordersharestockordergetAPIResponseModel
}

// AlibabawdkordersharestockordergetAPIResponseModel is 猫超商户订单拉取 成功返回结果
type AlibabawdkordersharestockordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *MaochaoOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
