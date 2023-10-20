package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenCombineitemSynchronize 组合商品接口
// taobao.qimen.combineitem.synchronize
//
// ERP调用奇门的接口,将商品信息同步给WMS
func TaobaoQimenCombineitemSynchronize(clt *core.SDKClient, req *qimen.TaobaoQimenCombineitemSynchronizeAPIRequest, resp *qimen.TaobaoQimenCombineitemSynchronizeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
