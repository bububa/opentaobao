package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaordcaligeniusrefundscheck 退款信息审核
// taobao.rdc.aligenius.refunds.check
//
// 根据退款信息，对退款单进行审核
func Taobaordcaligeniusrefundscheck(clt *core.SDKClient, req *util.TaobaordcaligeniusrefundscheckAPIRequest, session string) (*util.TaobaordcaligeniusrefundscheckAPIResponse, error) {
	var resp util.TaobaordcaligeniusrefundscheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
