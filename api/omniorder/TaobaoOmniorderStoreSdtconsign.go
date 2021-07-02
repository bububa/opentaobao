package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreSdtconsign 通知菜鸟裹裹发货
// taobao.omniorder.store.sdtconsign
//
// ISV取完单号后通知菜鸟裹裹发货
func TaobaoOmniorderStoreSdtconsign(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSdtconsignAPIRequest, session string) (*omniorder.TaobaoOmniorderStoreSdtconsignAPIResponse, error) {
	var resp omniorder.TaobaoOmniorderStoreSdtconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
