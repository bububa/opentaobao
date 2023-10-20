package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrfulfilllogisticsconsign 同城配门店备货通知
// tmall.nr.fulfill.logistics.consign
//
// 同城配业务备货通知，商家告诉平台门店的货已经准备好，可以发货了；
func Tmallnrfulfilllogisticsconsign(clt *core.SDKClient, req *tmallnr.TmallnrfulfilllogisticsconsignAPIRequest, session string) (*tmallnr.TmallnrfulfilllogisticsconsignAPIResponse, error) {
	var resp tmallnr.TmallnrfulfilllogisticsconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
