package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenpresalespackageconsign 预售预包尾款推单发货
// taobao.qimen.presalespackage.consign
//
// 预售预包尾款推单发货
func Taobaoqimenpresalespackageconsign(clt *core.SDKClient, req *qimen.TaobaoqimenpresalespackageconsignAPIRequest, session string) (*qimen.TaobaoqimenpresalespackageconsignAPIResponse, error) {
	var resp qimen.TaobaoqimenpresalespackageconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
