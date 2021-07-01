package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreDeliverconfigGetAPIRequest
查询门店发货配置内容 API请求
taobao.omniorder.store.deliverconfig.get

查询门店发货配置内容 */
type TaobaoOmniorderStoreDeliverconfigGetAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 是否是活动期
	_activity bool
}

// New
