package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrGetbyentidAPIResponse
多融通过企业ID得到一个企业的详细信息 API返回值
alibaba.alihealth.drug.kyt.dr.getbyentid

根据企业主键查看企业详细信息 */
type AlibabaAlihealthDrugKytDrGetbyentidAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrGetbyentidAPIResponseModel
}

// AlibabaAlihealthDrugKytDrGetbyentidAPIResponseModel is 多融通过企业ID得到一个企业的详细信息 成功返回结果
type AlibabaAlihealthDrugKytDrGetbyentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_getbyentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytDrGetbyentidResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
