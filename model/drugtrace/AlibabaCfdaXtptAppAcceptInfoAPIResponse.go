package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacfdaxtptappacceptinfoAPIResponse 协同平台数据下行接口 API返回值
// alibaba.cfda.xtpt.app.accept.info
//
// 协同平台数据下行接口
type AlibabacfdaxtptappacceptinfoAPIResponse struct {
	model.CommonResponse
	AlibabacfdaxtptappacceptinfoAPIResponseModel
}

// AlibabacfdaxtptappacceptinfoAPIResponseModel is 协同平台数据下行接口 成功返回结果
type AlibabacfdaxtptappacceptinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cfda_xtpt_app_accept_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabacfdaxtptappacceptinfoResult `json:"result,omitempty" xml:"result,omitempty"`
}
