package shop

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaomcnshopcatslistgetAPIResponse 店铺类目清单 API返回值
// taobao.mcn.shopcats.list.get
//
// 无需授权; 获取前台展示的店铺类目;
type TaobaomcnshopcatslistgetAPIResponse struct {
	model.CommonResponse
	TaobaomcnshopcatslistgetAPIResponseModel
}

// TaobaomcnshopcatslistgetAPIResponseModel is 店铺类目清单 成功返回结果
type TaobaomcnshopcatslistgetAPIResponseModel struct {
	XMLName xml.Name `xml:"mcn_shopcats_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 店铺类目列表信息
	ShopCats []ShopCat `json:"shop_cats,omitempty" xml:"shop_cats>shop_cat,omitempty"`
}
