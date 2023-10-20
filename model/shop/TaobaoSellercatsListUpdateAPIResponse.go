package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercatsListUpdateAPIResponse 更新卖家自定义类目 API返回值
// taobao.sellercats.list.update
//
// 此API更新卖家店铺内自定义类目 &lt;br/&gt;注：因为缓存的关系，添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
type TaobaoSellercatsListUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSellercatsListUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSellercatsListUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSellercatsListUpdateAPIResponseModel).Reset()
}

// TaobaoSellercatsListUpdateAPIResponseModel is 更新卖家自定义类目 成功返回结果
type TaobaoSellercatsListUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"sellercats_list_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回sellercat数据结构中的：cid,modified
	SellerCat *SellerCat `json:"seller_cat,omitempty" xml:"seller_cat,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSellercatsListUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SellerCat = nil
}

var poolTaobaoSellercatsListUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSellercatsListUpdateAPIResponse)
	},
}

// GetTaobaoSellercatsListUpdateAPIResponse 从 sync.Pool 获取 TaobaoSellercatsListUpdateAPIResponse
func GetTaobaoSellercatsListUpdateAPIResponse() *TaobaoSellercatsListUpdateAPIResponse {
	return poolTaobaoSellercatsListUpdateAPIResponse.Get().(*TaobaoSellercatsListUpdateAPIResponse)
}

// ReleaseTaobaoSellercatsListUpdateAPIResponse 将 TaobaoSellercatsListUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSellercatsListUpdateAPIResponse(v *TaobaoSellercatsListUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSellercatsListUpdateAPIResponse.Put(v)
}
