package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenItemsSynchronizeAPIRequest
商品同步接口 (批量) API请求
taobao.qimen.items.synchronize

ERP调用奇门的接口,批量同步商品信息给WMS */
type TaobaoQimenItemsSynchronizeAPIRequest struct {
	model.Params
	//
	_request *ItemsSynchronizeRequest
}

// New
