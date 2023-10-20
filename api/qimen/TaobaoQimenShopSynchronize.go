package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenshopsynchronize 店铺同步接口
// taobao.qimen.shop.synchronize
//
// 店铺同步接口描述
func Taobaoqimenshopsynchronize(clt *core.SDKClient, req *qimen.TaobaoqimenshopsynchronizeAPIRequest, session string) (*qimen.TaobaoqimenshopsynchronizeAPIResponse, error) {
	var resp qimen.TaobaoqimenshopsynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
