package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetbyrefentidAPIResponse 根据企业唯一标识查看企业详细信息 API返回值
// alibaba.alihealth.drug.kyt.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaAlihealthDrugKytGetbyrefentidAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytGetbyrefentidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytGetbyrefentidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytGetbyrefentidAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytGetbyrefentidAPIResponseModel is 根据企业唯一标识查看企业详细信息 成功返回结果
type AlibabaAlihealthDrugKytGetbyrefentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getbyrefentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytGetbyrefentidResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytGetbyrefentidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytGetbyrefentidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetbyrefentidAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytGetbyrefentidAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytGetbyrefentidAPIResponse
func GetAlibabaAlihealthDrugKytGetbyrefentidAPIResponse() *AlibabaAlihealthDrugKytGetbyrefentidAPIResponse {
	return poolAlibabaAlihealthDrugKytGetbyrefentidAPIResponse.Get().(*AlibabaAlihealthDrugKytGetbyrefentidAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytGetbyrefentidAPIResponse 将 AlibabaAlihealthDrugKytGetbyrefentidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytGetbyrefentidAPIResponse(v *AlibabaAlihealthDrugKytGetbyrefentidAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytGetbyrefentidAPIResponse.Put(v)
}
