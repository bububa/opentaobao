package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenItemsSynchronize 商品同步接口 (批量)
// taobao.qimen.items.synchronize
//
// ERP调用奇门的接口,批量同步商品信息给WMS
func TaobaoQimenItemsSynchronize(clt *core.SDKClient, req *qimen.TaobaoQimenItemsSynchronizeAPIRequest, resp *qimen.TaobaoQimenItemsSynchronizeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
