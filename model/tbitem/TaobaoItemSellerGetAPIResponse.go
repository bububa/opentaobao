package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemSellerGetAPIResponse 获取单个商品详细信息 API返回值
// taobao.item.seller.get
//
// 获取单个商品的全部信息
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoItemSellerGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemSellerGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemSellerGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemSellerGetAPIResponseModel).Reset()
}

// TaobaoItemSellerGetAPIResponseModel is 获取单个商品详细信息 成功返回结果
type TaobaoItemSellerGetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_seller_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品详细信息
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemSellerGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Item = nil
}

var poolTaobaoItemSellerGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemSellerGetAPIResponse)
	},
}

// GetTaobaoItemSellerGetAPIResponse 从 sync.Pool 获取 TaobaoItemSellerGetAPIResponse
func GetTaobaoItemSellerGetAPIResponse() *TaobaoItemSellerGetAPIResponse {
	return poolTaobaoItemSellerGetAPIResponse.Get().(*TaobaoItemSellerGetAPIResponse)
}

// ReleaseTaobaoItemSellerGetAPIResponse 将 TaobaoItemSellerGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemSellerGetAPIResponse(v *TaobaoItemSellerGetAPIResponse) {
	v.Reset()
	poolTaobaoItemSellerGetAPIResponse.Put(v)
}
