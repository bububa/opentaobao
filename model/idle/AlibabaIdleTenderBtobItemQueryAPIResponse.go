package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderBtobItemQueryAPIResponse 暗拍b2b商品查询 API返回值
// alibaba.idle.tender.btob.item.query
//
// 暗拍b2b商品查询
type AlibabaIdleTenderBtobItemQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTenderBtobItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleTenderBtobItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleTenderBtobItemQueryAPIResponseModel).Reset()
}

// AlibabaIdleTenderBtobItemQueryAPIResponseModel is 暗拍b2b商品查询 成功返回结果
type AlibabaIdleTenderBtobItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_btob_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleTenderBtobItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleTenderBtobItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTenderBtobItemQueryAPIResponse)
	},
}

// GetAlibabaIdleTenderBtobItemQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleTenderBtobItemQueryAPIResponse
func GetAlibabaIdleTenderBtobItemQueryAPIResponse() *AlibabaIdleTenderBtobItemQueryAPIResponse {
	return poolAlibabaIdleTenderBtobItemQueryAPIResponse.Get().(*AlibabaIdleTenderBtobItemQueryAPIResponse)
}

// ReleaseAlibabaIdleTenderBtobItemQueryAPIResponse 将 AlibabaIdleTenderBtobItemQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleTenderBtobItemQueryAPIResponse(v *AlibabaIdleTenderBtobItemQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleTenderBtobItemQueryAPIResponse.Put(v)
}
