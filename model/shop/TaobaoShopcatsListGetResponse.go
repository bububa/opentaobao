package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取前台展示的店铺类目 APIResponse
taobao.shopcats.list.get

获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
*/
type TaobaoShopcatsListGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"shopcats_list_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 店铺类目列表信息
    
    ShopCats   []ShopCat `json:"shop_cats,omitempty" xml:"