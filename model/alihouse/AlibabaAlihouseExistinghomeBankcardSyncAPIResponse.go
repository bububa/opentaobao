package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBankcardSyncAPIResponse 结算账号同步 API返回值
// alibaba.alihouse.existinghome.bankcard.sync
//
// 结算账号同步
type AlibabaAlihouseExistinghomeBankcardSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeBankcardSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBankcardSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeBankcardSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeBankcardSyncAPIResponseModel is 结算账号同步 成功返回结果
type AlibabaAlihouseExistinghomeBankcardSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_bankcard_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeBankcardSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBankcardSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeBankcardSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBankcardSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeBankcardSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeBankcardSyncAPIResponse
func GetAlibabaAlihouseExistinghomeBankcardSyncAPIResponse() *AlibabaAlihouseExistinghomeBankcardSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeBankcardSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeBankcardSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeBankcardSyncAPIResponse 将 AlibabaAlihouseExistinghomeBankcardSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeBankcardSyncAPIResponse(v *AlibabaAlihouseExistinghomeBankcardSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeBankcardSyncAPIResponse.Put(v)
}
