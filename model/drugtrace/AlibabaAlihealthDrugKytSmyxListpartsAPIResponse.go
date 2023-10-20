package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytsmyxlistpartsAPIResponse 药店查询往来单位 API返回值
// alibaba.alihealth.drug.kyt.smyx.listparts
//
// 查询往来单位列表
type AlibabaalihealthdrugkytsmyxlistpartsAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytsmyxlistpartsAPIResponseModel
}

// AlibabaalihealthdrugkytsmyxlistpartsAPIResponseModel is 药店查询往来单位 成功返回结果
type AlibabaalihealthdrugkytsmyxlistpartsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_smyx_listparts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytsmyxlistpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
