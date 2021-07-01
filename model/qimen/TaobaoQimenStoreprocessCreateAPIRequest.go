package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreprocessCreateAPIRequest
仓内加工单创建接口 API请求
taobao.qimen.storeprocess.create

ERP调用奇门的接口,创建仓内加工单 */
type TaobaoQimenStoreprocessCreateAPIRequest struct {
	model.Params
	//
	_request *StoreProcessCreateRequest
}

// New
