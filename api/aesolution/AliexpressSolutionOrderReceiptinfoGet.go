package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionorderreceiptinfoget Get Order Receipt Info
// aliexpress.solution.order.receiptinfo.get
//
// Get Order Receipt Info, Support multi stores requirements for Turkey sellers.
func Aliexpresssolutionorderreceiptinfoget(clt *core.SDKClient, req *aesolution.AliexpresssolutionorderreceiptinfogetAPIRequest, session string) (*aesolution.AliexpresssolutionorderreceiptinfogetAPIResponse, error) {
	var resp aesolution.AliexpresssolutionorderreceiptinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
