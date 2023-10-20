package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse 根据企业唯一标识查看企业详细信息 API返回值
// alibaba.alihealth.drug.kyt.wes.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponseModel is 根据企业唯一标识查看企业详细信息 成功返回结果
type AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_getbyrefentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesGetbyrefentidResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse
func GetAlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse() *AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse {
	return poolAlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse.Get().(*AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse 将 AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse(v *AlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesGetbyrefentidAPIResponse.Put(v)
}
