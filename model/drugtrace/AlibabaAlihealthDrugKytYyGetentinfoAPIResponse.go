package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytyygetentinfoAPIResponse 得到企业信息 API返回值
// alibaba.alihealth.drug.kyt.yy.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaalihealthdrugkytyygetentinfoAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytyygetentinfoAPIResponseModel
}

// AlibabaalihealthdrugkytyygetentinfoAPIResponseModel is 得到企业信息 成功返回结果
type AlibabaalihealthdrugkytyygetentinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_yy_getentinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytyygetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
