package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse 进件接口 API返回值
// alibaba.alihouse.existinghome.companyinfo.submit
//
// 进件接口
type AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponseModel is 进件接口 成功返回结果
type AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_companyinfo_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeCompanyinfoSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse
func GetAlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse() *AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse {
	return poolAlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse.Get().(*AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse 将 AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse(v *AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeCompanyinfoSubmitAPIResponse.Put(v)
}
