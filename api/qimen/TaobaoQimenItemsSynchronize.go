package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenitemssynchronize 商品同步接口 (批量)
// taobao.qimen.items.synchronize
//
// ERP调用奇门的接口,批量同步商品信息给WMS
func Taobaoqimenitemssynchronize(clt *core.SDKClient, req *qimen.TaobaoqimenitemssynchronizeAPIRequest, session string) (*qimen.TaobaoqimenitemssynchronizeAPIResponse, error) {
	var resp qimen.TaobaoqimenitemssynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
