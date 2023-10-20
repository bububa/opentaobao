package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytListauthsAPIResponse 企业搜索自己授权的物流企业 API返回值
// alibaba.alihealth.drug.kyt.listauths
//
// 企业搜索自己授权的物流企业
type AlibabaAlihealthDrugKytListauthsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytListauthsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytListauthsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytListauthsAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytListauthsAPIResponseModel is 企业搜索自己授权的物流企业 成功返回结果
type AlibabaAlihealthDrugKytListauthsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_listauths_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytListauthsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytListauthsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytListauthsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytListauthsAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytListauthsAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytListauthsAPIResponse
func GetAlibabaAlihealthDrugKytListauthsAPIResponse() *AlibabaAlihealthDrugKytListauthsAPIResponse {
	return poolAlibabaAlihealthDrugKytListauthsAPIResponse.Get().(*AlibabaAlihealthDrugKytListauthsAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytListauthsAPIResponse 将 AlibabaAlihealthDrugKytListauthsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytListauthsAPIResponse(v *AlibabaAlihealthDrugKytListauthsAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytListauthsAPIResponse.Put(v)
}
