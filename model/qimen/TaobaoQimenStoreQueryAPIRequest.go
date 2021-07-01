package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreQueryAPIRequest
门店信息查询接口 API请求
taobao.qimen.store.query

商家在ERP等系统中调用该接口，查询门店相关信息 */
type TaobaoQimenStoreQueryAPIRequest struct {
	model.Params
	// 已分配的线上门店ID
	_storeId int64
}

// New
