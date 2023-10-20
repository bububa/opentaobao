package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvItemEditAPIResponse 服务商闲鱼商品编辑 API返回值
// alibaba.idle.isv.item.edit
//
// 服务商ISV闲鱼商品编辑操作
type AlibabaIdleIsvItemEditAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvItemEditAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvItemEditAPIResponseModel).Reset()
}

// AlibabaIdleIsvItemEditAPIResponseModel is 服务商闲鱼商品编辑 成功返回结果
type AlibabaIdleIsvItemEditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_item_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果result
	Result *AlibabaIdleIsvItemEditTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvItemEditAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvItemEditAPIResponse)
	},
}

// GetAlibabaIdleIsvItemEditAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvItemEditAPIResponse
func GetAlibabaIdleIsvItemEditAPIResponse() *AlibabaIdleIsvItemEditAPIResponse {
	return poolAlibabaIdleIsvItemEditAPIResponse.Get().(*AlibabaIdleIsvItemEditAPIResponse)
}

// ReleaseAlibabaIdleIsvItemEditAPIResponse 将 AlibabaIdleIsvItemEditAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvItemEditAPIResponse(v *AlibabaIdleIsvItemEditAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvItemEditAPIResponse.Put(v)
}
