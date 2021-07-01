package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreCollectconfigUpdateAPIRequest
门店自提配置修改 API请求
taobao.omniorder.store.collectconfig.update

修改门店自提配置内容 */
type TaobaoOmniorderStoreCollectconfigUpdateAPIRequest struct {
	model.Params
	// 门店自提配置
	_storeCollectConfig *StoreCollectConfig
	// 门店ID
	_storeId int64
}

// New
