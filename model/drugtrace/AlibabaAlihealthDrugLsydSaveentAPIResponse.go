package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdruglsydsaveentAPIResponse 零售药店往来单位新增 API返回值
// alibaba.alihealth.drug.lsyd.saveent
//
// 新增往来单位企业记录
type AlibabaalihealthdruglsydsaveentAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdruglsydsaveentAPIResponseModel
}

// AlibabaalihealthdruglsydsaveentAPIResponseModel is 零售药店往来单位新增 成功返回结果
type AlibabaalihealthdruglsydsaveentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_lsyd_saveent_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 往来单位新增接口返回
	Result *AlibabaalihealthdruglsydsaveentResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
