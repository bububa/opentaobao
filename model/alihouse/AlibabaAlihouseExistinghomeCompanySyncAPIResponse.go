package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeCompanySyncAPIResponse 二手房公司同步接口 API返回值
// alibaba.alihouse.existinghome.company.sync
//
// 二手房公司同步接口
type AlibabaAlihouseExistinghomeCompanySyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeCompanySyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeCompanySyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeCompanySyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeCompanySyncAPIResponseModel is 二手房公司同步接口 成功返回结果
type AlibabaAlihouseExistinghomeCompanySyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_company_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeCompanySyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeCompanySyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeCompanySyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeCompanySyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeCompanySyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeCompanySyncAPIResponse
func GetAlibabaAlihouseExistinghomeCompanySyncAPIResponse() *AlibabaAlihouseExistinghomeCompanySyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeCompanySyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeCompanySyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeCompanySyncAPIResponse 将 AlibabaAlihouseExistinghomeCompanySyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeCompanySyncAPIResponse(v *AlibabaAlihouseExistinghomeCompanySyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeCompanySyncAPIResponse.Put(v)
}
