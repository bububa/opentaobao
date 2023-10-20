package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemSkuGetAPIResponse 获取全渠道门店商品sku API返回值
// taobao.omniitem.sku.get
//
// 通过skuId或者skuOutId查询全渠道门店商品sku信息
type TaobaoOmniitemSkuGetAPIResponse struct {
	model.CommonResponse
	TaobaoOmniitemSkuGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniitemSkuGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniitemSkuGetAPIResponseModel).Reset()
}

// TaobaoOmniitemSkuGetAPIResponseModel is 获取全渠道门店商品sku 成功返回结果
type TaobaoOmniitemSkuGetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_sku_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OmniResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniitemSkuGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniitemSkuGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniitemSkuGetAPIResponse)
	},
}

// GetTaobaoOmniitemSkuGetAPIResponse 从 sync.Pool 获取 TaobaoOmniitemSkuGetAPIResponse
func GetTaobaoOmniitemSkuGetAPIResponse() *TaobaoOmniitemSkuGetAPIResponse {
	return poolTaobaoOmniitemSkuGetAPIResponse.Get().(*TaobaoOmniitemSkuGetAPIResponse)
}

// ReleaseTaobaoOmniitemSkuGetAPIResponse 将 TaobaoOmniitemSkuGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniitemSkuGetAPIResponse(v *TaobaoOmniitemSkuGetAPIResponse) {
	v.Reset()
	poolTaobaoOmniitemSkuGetAPIResponse.Put(v)
}
