package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest
switchstatus.update API请求
taobao.omniorder.store.switchstatus.update

变更门店发货、门店自提状态 */
type TaobaoOmniorderStoreSwitchstatusUpdateAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 门店发货自提状态
	_status string
}

// New
