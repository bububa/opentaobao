package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkordersharestockfulfillgetAPIResponse 商户订单履约数据获取 API返回值
// alibaba.wdkorder.sharestock.fulfill.get
//
// 商户订单履约数据获取
type AlibabawdkordersharestockfulfillgetAPIResponse struct {
	model.CommonResponse
	AlibabawdkordersharestockfulfillgetAPIResponseModel
}

// AlibabawdkordersharestockfulfillgetAPIResponseModel is 商户订单履约数据获取 成功返回结果
type AlibabawdkordersharestockfulfillgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_fulfill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *MaochaoOrderFulfillQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
