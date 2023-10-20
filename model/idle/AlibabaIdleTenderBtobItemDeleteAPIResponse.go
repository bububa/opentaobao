package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderBtobItemDeleteAPIResponse 暗拍b2b商品下架/删除 API返回值
// alibaba.idle.tender.btob.item.delete
//
// 暗拍b2b商品下架/删除
type AlibabaIdleTenderBtobItemDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTenderBtobItemDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleTenderBtobItemDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleTenderBtobItemDeleteAPIResponseModel).Reset()
}

// AlibabaIdleTenderBtobItemDeleteAPIResponseModel is 暗拍b2b商品下架/删除 成功返回结果
type AlibabaIdleTenderBtobItemDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_btob_item_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleTenderBtobItemDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleTenderBtobItemDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTenderBtobItemDeleteAPIResponse)
	},
}

// GetAlibabaIdleTenderBtobItemDeleteAPIResponse 从 sync.Pool 获取 AlibabaIdleTenderBtobItemDeleteAPIResponse
func GetAlibabaIdleTenderBtobItemDeleteAPIResponse() *AlibabaIdleTenderBtobItemDeleteAPIResponse {
	return poolAlibabaIdleTenderBtobItemDeleteAPIResponse.Get().(*AlibabaIdleTenderBtobItemDeleteAPIResponse)
}

// ReleaseAlibabaIdleTenderBtobItemDeleteAPIResponse 将 AlibabaIdleTenderBtobItemDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleTenderBtobItemDeleteAPIResponse(v *AlibabaIdleTenderBtobItemDeleteAPIResponse) {
	v.Reset()
	poolAlibabaIdleTenderBtobItemDeleteAPIResponse.Put(v)
}
