package shop

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercatsListGetAPIResponse 获取前台展示的店铺内卖家自定义商品类目 API返回值
// taobao.sellercats.list.get
//
// 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
type TaobaoSellercatsListGetAPIResponse struct {
	model.CommonResponse
	TaobaoSellercatsListGetAPIResponseModel
}

// TaobaoSellercatsListGetAPIResponseModel is 获取前台展示的店铺内卖家自定义商品类目 成功返回结果
type TaobaoSellercatsListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"sellercats_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 卖家自定义类目
	SellerCats []SellerCat `json:"seller_cats,omitempty" xml:"seller_cats>seller_cat,omitempty"`
}
