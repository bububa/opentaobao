package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeQuotationSyncAPIResponse 二手房行情数据同步 API返回值
// alibaba.alihouse.existinghome.quotation.sync
//
// 二手房行情数据同步
type AlibabaAlihouseExistinghomeQuotationSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeQuotationSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeQuotationSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeQuotationSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeQuotationSyncAPIResponseModel is 二手房行情数据同步 成功返回结果
type AlibabaAlihouseExistinghomeQuotationSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_quotation_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihouseExistinghomeQuotationSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeQuotationSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeQuotationSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeQuotationSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeQuotationSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeQuotationSyncAPIResponse
func GetAlibabaAlihouseExistinghomeQuotationSyncAPIResponse() *AlibabaAlihouseExistinghomeQuotationSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeQuotationSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeQuotationSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeQuotationSyncAPIResponse 将 AlibabaAlihouseExistinghomeQuotationSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeQuotationSyncAPIResponse(v *AlibabaAlihouseExistinghomeQuotationSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeQuotationSyncAPIResponse.Put(v)
}
