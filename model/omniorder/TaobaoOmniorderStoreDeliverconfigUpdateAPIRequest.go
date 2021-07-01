package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest
修改门店发货配置内容 API请求
taobao.omniorder.store.deliverconfig.update

修改门店发货配置内容 */
type TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 卖家发货配置
	_storeDeliverConfig *StoreDeliverConfig
}

// New
