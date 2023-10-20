package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenitemlackreport 发货单缺货通知接口
// taobao.qimen.itemlack.report
//
// WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
func Taobaoqimenitemlackreport(clt *core.SDKClient, req *qimen.TaobaoqimenitemlackreportAPIRequest, session string) (*qimen.TaobaoqimenitemlackreportAPIResponse, error) {
	var resp qimen.TaobaoqimenitemlackreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
