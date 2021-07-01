package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreprocessConfirmAPIRequest
仓内加工单确认接口 API请求
taobao.qimen.storeprocess.confirm

WMS调用奇门的接口,回传仓内加工单创建情况 */
type TaobaoQimenStoreprocessConfirmAPIRequest struct {
	model.Params
	//
	_request *StoreProcessConfirmRequest
}

// New
