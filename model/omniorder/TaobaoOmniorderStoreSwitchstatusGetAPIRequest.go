package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreSwitchstatusGetAPIRequest
switchstatus.get API请求
taobao.omniorder.store.switchstatus.get

查询门店发货、门店自提状态 */
type TaobaoOmniorderStoreSwitchstatusGetAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 卖家ID
	_sellerId int64
}

// New
