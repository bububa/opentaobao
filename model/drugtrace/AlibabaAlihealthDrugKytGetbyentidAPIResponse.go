package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetbyentidAPIResponse 根据企业主键查看企业详细信息 API返回值
// alibaba.alihealth.drug.kyt.getbyentid
//
// 根据企业主键查看企业详细信息
type AlibabaAlihealthDrugKytGetbyentidAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytGetbyentidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytGetbyentidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytGetbyentidAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytGetbyentidAPIResponseModel is 根据企业主键查看企业详细信息 成功返回结果
type AlibabaAlihealthDrugKytGetbyentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getbyentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytGetbyentidResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytGetbyentidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytGetbyentidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetbyentidAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytGetbyentidAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytGetbyentidAPIResponse
func GetAlibabaAlihealthDrugKytGetbyentidAPIResponse() *AlibabaAlihealthDrugKytGetbyentidAPIResponse {
	return poolAlibabaAlihealthDrugKytGetbyentidAPIResponse.Get().(*AlibabaAlihealthDrugKytGetbyentidAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytGetbyentidAPIResponse 将 AlibabaAlihealthDrugKytGetbyentidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytGetbyentidAPIResponse(v *AlibabaAlihealthDrugKytGetbyentidAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytGetbyentidAPIResponse.Put(v)
}
