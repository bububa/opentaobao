package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreitemQueryAPIRequest
门店关联商品查询接口 API请求
taobao.qimen.storeitem.query

商家在ERP等系统中调用该接口，查询某门店所关联的线上商品列表 */
type TaobaoQimenStoreitemQueryAPIRequest struct {
	model.Params
	// 当前页面
	_page int64
	// 线上门店id
	_storeId int64
}

// New
