package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicepayoutgetAPIResponse 获取赔付计时列表数据 API返回值
// alibaba.einvoice.payout.get
//
// 获取赔付计时列表数据
type AlibabaeinvoicepayoutgetAPIResponse struct {
	model.CommonResponse
	AlibabaeinvoicepayoutgetAPIResponseModel
}

// AlibabaeinvoicepayoutgetAPIResponseModel is 获取赔付计时列表数据 成功返回结果
type AlibabaeinvoicepayoutgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_payout_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *OrderRightsResult `json:"result,omitempty" xml:"result,omitempty"`
	// 查询结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
