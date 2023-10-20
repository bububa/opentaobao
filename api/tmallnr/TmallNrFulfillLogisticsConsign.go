package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrFulfillLogisticsConsign 同城配门店备货通知
// tmall.nr.fulfill.logistics.consign
//
// 同城配业务备货通知，商家告诉平台门店的货已经准备好，可以发货了；
func TmallNrFulfillLogisticsConsign(clt *core.SDKClient, req *tmallnr.TmallNrFulfillLogisticsConsignAPIRequest, resp *tmallnr.TmallNrFulfillLogisticsConsignAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
