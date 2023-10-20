package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenstoreprocessconfirm 仓内加工单确认接口
// taobao.qimen.storeprocess.confirm
//
// WMS调用奇门的接口,回传仓内加工单创建情况
func Taobaoqimenstoreprocessconfirm(clt *core.SDKClient, req *qimen.TaobaoqimenstoreprocessconfirmAPIRequest, session string) (*qimen.TaobaoqimenstoreprocessconfirmAPIResponse, error) {
	var resp qimen.TaobaoqimenstoreprocessconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
