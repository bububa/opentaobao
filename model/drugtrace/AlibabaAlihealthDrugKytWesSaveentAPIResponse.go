package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytwessaveentAPIResponse 新增往来单位企业记录 API返回值
// alibaba.alihealth.drug.kyt.wes.saveent
//
// 新增往来单位企业记录
type AlibabaalihealthdrugkytwessaveentAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytwessaveentAPIResponseModel
}

// AlibabaalihealthdrugkytwessaveentAPIResponseModel is 新增往来单位企业记录 成功返回结果
type AlibabaalihealthdrugkytwessaveentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_saveent_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 往来单位新增接口返回
	Result *AlibabaalihealthdrugkytwessaveentResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
