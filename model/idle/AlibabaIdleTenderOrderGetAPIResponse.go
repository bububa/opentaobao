package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderordergetAPIResponse 暗拍读取订单 API返回值
// alibaba.idle.tender.order.get
//
// 查询省心卖暗拍项目订单
type AlibabaidletenderordergetAPIResponse struct {
	model.CommonResponse
	AlibabaidletenderordergetAPIResponseModel
}

// AlibabaidletenderordergetAPIResponseModel is 暗拍读取订单 成功返回结果
type AlibabaidletenderordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *AlibabaidletenderordergetResult `json:"result,omitempty" xml:"result,omitempty"`
}
