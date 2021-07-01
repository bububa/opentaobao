package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenItemstoreQueryAPIRequest
商品关联门店查询接口 API请求
taobao.qimen.itemstore.query

商家在ERP等系统中调用该接口，查询线上商品所关联的门店列表 */
type TaobaoQimenItemstoreQueryAPIRequest struct {
	model.Params
	// 当前查询的页面编码
	_page int64
	// 线上商品ID
	_itemId int64
}

// New
