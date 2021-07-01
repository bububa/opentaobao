package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoShopcatsListGetAPIRequest
获取前台展示的店铺类目 API请求
taobao.shopcats.list.get

获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异） */
type TaobaoShopcatsListGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
	_fields []string
}

// New
