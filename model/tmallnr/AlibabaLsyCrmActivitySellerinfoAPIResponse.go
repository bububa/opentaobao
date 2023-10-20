package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmactivitysellerinfoAPIResponse 商家信息推送 API返回值
// alibaba.lsy.crm.activity.sellerinfo
//
// 本地团商家信息推送
type AlibabalsycrmactivitysellerinfoAPIResponse struct {
	model.CommonResponse
	AlibabalsycrmactivitysellerinfoAPIResponseModel
}

// AlibabalsycrmactivitysellerinfoAPIResponseModel is 商家信息推送 成功返回结果
type AlibabalsycrmactivitysellerinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_sellerinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
