package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreCollectconfigGetAPIRequest
查询门店自提配置内容 API请求
taobao.omniorder.store.collectconfig.get

查询门店自提配置内容 */
type TaobaoOmniorderStoreCollectconfigGetAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 是否是活动期
	_activity bool
}

// New
