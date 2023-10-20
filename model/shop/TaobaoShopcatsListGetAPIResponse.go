package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoShopcatsListGetAPIResponse 获取前台展示的店铺类目 API返回值
// taobao.shopcats.list.get
//
// 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
type TaobaoShopcatsListGetAPIResponse struct {
	model.CommonResponse
	TaobaoShopcatsListGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoShopcatsListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoShopcatsListGetAPIResponseModel).Reset()
}

// TaobaoShopcatsListGetAPIResponseModel is 获取前台展示的店铺类目 成功返回结果
type TaobaoShopcatsListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"shopcats_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 店铺类目列表信息
	ShopCats []ShopCat `json:"shop_cats,omitempty" xml:"shop_cats>shop_cat,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoShopcatsListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ShopCats = m.ShopCats[:0]
}

var poolTaobaoShopcatsListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoShopcatsListGetAPIResponse)
	},
}

// GetTaobaoShopcatsListGetAPIResponse 从 sync.Pool 获取 TaobaoShopcatsListGetAPIResponse
func GetTaobaoShopcatsListGetAPIResponse() *TaobaoShopcatsListGetAPIResponse {
	return poolTaobaoShopcatsListGetAPIResponse.Get().(*TaobaoShopcatsListGetAPIResponse)
}

// ReleaseTaobaoShopcatsListGetAPIResponse 将 TaobaoShopcatsListGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoShopcatsListGetAPIResponse(v *TaobaoShopcatsListGetAPIResponse) {
	v.Reset()
	poolTaobaoShopcatsListGetAPIResponse.Put(v)
}
