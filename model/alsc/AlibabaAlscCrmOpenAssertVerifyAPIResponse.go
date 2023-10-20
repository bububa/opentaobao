package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmopenassertverifyAPIResponse 资产核销接口 API返回值
// alibaba.alsc.crm.open.assert.verify
//
// 核销储值，积分，券资产
type AlibabaalsccrmopenassertverifyAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmopenassertverifyAPIResponseModel
}

// AlibabaalsccrmopenassertverifyAPIResponseModel is 资产核销接口 成功返回结果
type AlibabaalsccrmopenassertverifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_open_assert_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
