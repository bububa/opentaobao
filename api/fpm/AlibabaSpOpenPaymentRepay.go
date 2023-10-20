package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// Alibabaspopenpaymentrepay 智付重新打款
// alibaba.sp.open.payment.repay
//
// 智付重新打款
func Alibabaspopenpaymentrepay(clt *core.SDKClient, req *fpm.AlibabaspopenpaymentrepayAPIRequest, session string) (*fpm.AlibabaspopenpaymentrepayAPIResponse, error) {
	var resp fpm.AlibabaspopenpaymentrepayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
