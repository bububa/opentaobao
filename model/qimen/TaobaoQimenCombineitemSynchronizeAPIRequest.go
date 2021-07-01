package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenCombineitemSynchronizeAPIRequest
组合商品接口 API请求
taobao.qimen.combineitem.synchronize

ERP调用奇门的接口,将商品信息同步给WMS */
type TaobaoQimenCombineitemSynchronizeAPIRequest struct {
	model.Params
	//
	_request *CombineItemSyncRequest
}

// New
