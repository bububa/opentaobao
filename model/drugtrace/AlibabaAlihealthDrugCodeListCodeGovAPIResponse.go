package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodelistcodegovAPIResponse 政府码查询接口 API返回值
// alibaba.alihealth.drug.code.list.code.gov
//
// 政府码查询接口
type AlibabaalihealthdrugcodelistcodegovAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugcodelistcodegovAPIResponseModel
}

// AlibabaalihealthdrugcodelistcodegovAPIResponseModel is 政府码查询接口 成功返回结果
type AlibabaalihealthdrugcodelistcodegovAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_list_code_gov_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaalihealthdrugcodelistcodegovResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
