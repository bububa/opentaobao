package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemItemPublishAPIResponse 全渠道门店商品轻发布 API返回值
// taobao.omniitem.item.publish
//
// 全渠道门店商品轻发布
type TaobaoOmniitemItemPublishAPIResponse struct {
	model.CommonResponse
	TaobaoOmniitemItemPublishAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniitemItemPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniitemItemPublishAPIResponseModel).Reset()
}

// TaobaoOmniitemItemPublishAPIResponseModel is 全渠道门店商品轻发布 成功返回结果
type TaobaoOmniitemItemPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_item_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ItemLightPublishResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniitemItemPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniitemItemPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniitemItemPublishAPIResponse)
	},
}

// GetTaobaoOmniitemItemPublishAPIResponse 从 sync.Pool 获取 TaobaoOmniitemItemPublishAPIResponse
func GetTaobaoOmniitemItemPublishAPIResponse() *TaobaoOmniitemItemPublishAPIResponse {
	return poolTaobaoOmniitemItemPublishAPIResponse.Get().(*TaobaoOmniitemItemPublishAPIResponse)
}

// ReleaseTaobaoOmniitemItemPublishAPIResponse 将 TaobaoOmniitemItemPublishAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniitemItemPublishAPIResponse(v *TaobaoOmniitemItemPublishAPIResponse) {
	v.Reset()
	poolTaobaoOmniitemItemPublishAPIResponse.Put(v)
}
