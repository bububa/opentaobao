package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoShopSellerGetAPIResponse 卖家店铺基础信息查询 API返回值
// taobao.shop.seller.get
//
// 获取卖家店铺的基本信息
type TaobaoShopSellerGetAPIResponse struct {
	model.CommonResponse
	TaobaoShopSellerGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoShopSellerGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoShopSellerGetAPIResponseModel).Reset()
}

// TaobaoShopSellerGetAPIResponseModel is 卖家店铺基础信息查询 成功返回结果
type TaobaoShopSellerGetAPIResponseModel struct {
	XMLName xml.Name `xml:"shop_seller_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 店铺信息
	Shop *Shop `json:"shop,omitempty" xml:"shop,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoShopSellerGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Shop = nil
}

var poolTaobaoShopSellerGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoShopSellerGetAPIResponse)
	},
}

// GetTaobaoShopSellerGetAPIResponse 从 sync.Pool 获取 TaobaoShopSellerGetAPIResponse
func GetTaobaoShopSellerGetAPIResponse() *TaobaoShopSellerGetAPIResponse {
	return poolTaobaoShopSellerGetAPIResponse.Get().(*TaobaoShopSellerGetAPIResponse)
}

// ReleaseTaobaoShopSellerGetAPIResponse 将 TaobaoShopSellerGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoShopSellerGetAPIResponse(v *TaobaoShopSellerGetAPIResponse) {
	v.Reset()
	poolTaobaoShopSellerGetAPIResponse.Put(v)
}
