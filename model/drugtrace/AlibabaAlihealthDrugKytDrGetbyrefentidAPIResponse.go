package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse 多融通过一个企业唯一标识查询企业详细信息 API返回值
// alibaba.alihealth.drug.kyt.dr.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponseModel
}

// AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponseModel is 多融通过一个企业唯一标识查询企业详细信息 成功返回结果
type AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_getbyrefentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytDrGetbyrefentidResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
