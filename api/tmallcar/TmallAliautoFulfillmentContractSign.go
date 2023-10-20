package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautofulfillmentcontractsign 合同签署
// tmall.aliauto.fulfillment.contract.sign
//
// 商家回传用户签署的合同信息
func Tmallaliautofulfillmentcontractsign(clt *core.SDKClient, req *tmallcar.TmallaliautofulfillmentcontractsignAPIRequest, session string) (*tmallcar.TmallaliautofulfillmentcontractsignAPIResponse, error) {
	var resp tmallcar.TmallaliautofulfillmentcontractsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
