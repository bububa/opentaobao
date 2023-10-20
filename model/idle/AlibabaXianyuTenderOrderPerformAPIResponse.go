package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaxianyutenderorderperformAPIResponse 闲鱼暗拍订单履约 API返回值
// alibaba.xianyu.tender.order.perform
//
// 闲鱼暗拍订单履约
type AlibabaxianyutenderorderperformAPIResponse struct {
	model.CommonResponse
	AlibabaxianyutenderorderperformAPIResponseModel
}

// AlibabaxianyutenderorderperformAPIResponseModel is 闲鱼暗拍订单履约 成功返回结果
type AlibabaxianyutenderorderperformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xianyu_tender_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaxianyutenderorderperformResult `json:"result,omitempty" xml:"result,omitempty"`
}
