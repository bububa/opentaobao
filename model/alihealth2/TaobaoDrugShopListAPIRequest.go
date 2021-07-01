package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDrugShopListAPIRequest
查询卖家外卖店列表 API请求
taobao.drug.shop.list

查询卖家外卖店列表 */
type TaobaoDrugShopListAPIRequest struct {
	model.Params
	// 查询关键字
	_keywords string
	// 店铺状态，歇业：0，营业：1，所有：-1
	_status int64
	// 页码
	_page int64
	// 每页条数
	_pageSize int64
}

// New
