package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取前台展示的店铺类目 API返回值 
taobao.shopcats.list.get

获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
*/
type TaobaoShopcatsListGetAPIResponse struct {
    model.CommonResponse
    TaobaoShopcatsListGetAPIResponseModel
}

// 获取前台展示的店铺类目 成功返回结果
type TaobaoShopcatsListGetAPIResponseModel struct {
    XMLName xml.Name `xml:"shopcats_list_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 店铺类目列表信息
    ShopCats   []ShopCat `json:"shop_cats,omitempty" xml:"shop_cats>shop_cat,omitempty"`
}
