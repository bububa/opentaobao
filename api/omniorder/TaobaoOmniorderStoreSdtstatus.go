package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

/* TaobaoOmniorderStoreSdtstatus
菜鸟裹裹运单状态查询
taobao.omniorder.store.sdtstatus

提供给商家查询运力单的状态。 */
func TaobaoOmniorderStoreSdtstatus(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSdtstatusAPIRequest, session string) (*omniorder.TaobaoOmniorderStoreSdtstatusAPIResponse, error) {
	var resp omniorder.TaobaoOmniorderStoreSdtstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
