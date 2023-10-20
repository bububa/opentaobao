package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse 码拼箱解除父子关系接口 API返回值
// alibaba.alihealth.drug.kyt.scqy.putpackageunbind
//
// 码拼箱解除父子关系接口
type AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponseModel is 码拼箱解除父子关系接口 成功返回结果
type AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_putpackageunbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugKytScqyPutpackageunbindResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse
func GetAlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse() *AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse {
	return poolAlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse.Get().(*AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse 将 AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse(v *AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse.Put(v)
}
