package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytscqycodeprocessAPIResponse 关联关系处理查询 API返回值
// alibaba.alihealth.drug.kyt.scqy.codeprocess
//
// 关联关系处理查询
type AlibabaalihealthdrugkytscqycodeprocessAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytscqycodeprocessAPIResponseModel
}

// AlibabaalihealthdrugkytscqycodeprocessAPIResponseModel is 关联关系处理查询 成功返回结果
type AlibabaalihealthdrugkytscqycodeprocessAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_codeprocess_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytscqycodeprocessResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
