package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreDeleteAPIRequest
门店删除接口 API请求
taobao.qimen.store.delete

商家在ERP等系统中调用该接口，删除线下门店 */
type TaobaoQimenStoreDeleteAPIRequest struct {
	model.Params
	// 要删除的门店id
	_storeId int64
}

// New
