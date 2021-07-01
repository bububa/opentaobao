package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenSingleitemSynchronizeAPIRequest
商品同步接口 API请求
taobao.qimen.singleitem.synchronize

ERP调用奇门的接口,同步商品信息给WMS */
type TaobaoQimenSingleitemSynchronizeAPIRequest struct {
	model.Params
	//
	_request *ItemSynRequest
}

// New
