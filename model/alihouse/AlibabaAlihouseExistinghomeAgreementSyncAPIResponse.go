package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeAgreementSyncAPIResponse 二手房电子协议数据同步 API返回值
// alibaba.alihouse.existinghome.agreement.sync
//
// 二手房电子协议数据同步
type AlibabaAlihouseExistinghomeAgreementSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeAgreementSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeAgreementSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeAgreementSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeAgreementSyncAPIResponseModel is 二手房电子协议数据同步 成功返回结果
type AlibabaAlihouseExistinghomeAgreementSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_agreement_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeAgreementSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeAgreementSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeAgreementSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeAgreementSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeAgreementSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeAgreementSyncAPIResponse
func GetAlibabaAlihouseExistinghomeAgreementSyncAPIResponse() *AlibabaAlihouseExistinghomeAgreementSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeAgreementSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeAgreementSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeAgreementSyncAPIResponse 将 AlibabaAlihouseExistinghomeAgreementSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeAgreementSyncAPIResponse(v *AlibabaAlihouseExistinghomeAgreementSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeAgreementSyncAPIResponse.Put(v)
}
