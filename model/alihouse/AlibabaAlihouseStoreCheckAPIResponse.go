package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseStoreCheckAPIResponse 门店对账查询工具 API返回值
// alibaba.alihouse.store.check
//
// 门店对账查询工具
type AlibabaAlihouseStoreCheckAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseStoreCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseStoreCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseStoreCheckAPIResponseModel).Reset()
}

// AlibabaAlihouseStoreCheckAPIResponseModel is 门店对账查询工具 成功返回结果
type AlibabaAlihouseStoreCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_store_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaAlihouseStoreCheckResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseStoreCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseStoreCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseStoreCheckAPIResponse)
	},
}

// GetAlibabaAlihouseStoreCheckAPIResponse 从 sync.Pool 获取 AlibabaAlihouseStoreCheckAPIResponse
func GetAlibabaAlihouseStoreCheckAPIResponse() *AlibabaAlihouseStoreCheckAPIResponse {
	return poolAlibabaAlihouseStoreCheckAPIResponse.Get().(*AlibabaAlihouseStoreCheckAPIResponse)
}

// ReleaseAlibabaAlihouseStoreCheckAPIResponse 将 AlibabaAlihouseStoreCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseStoreCheckAPIResponse(v *AlibabaAlihouseStoreCheckAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseStoreCheckAPIResponse.Put(v)
}
