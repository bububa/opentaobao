package drugtrace

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponseModel is 多融通过一个企业唯一标识查询企业详细信息 成功返回结果
type AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_getbyrefentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytDrGetbyrefentidResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse
func GetAlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse() *AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse {
	return poolAlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse.Get().(*AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse 将 AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse(v *AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse.Put(v)
}
