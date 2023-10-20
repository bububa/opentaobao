package fans

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fans"
)

// Tmallfanscashpoolcheckpay 检查资金池付款状态
// tmall.fans.cashpool.checkpay
//
// 检查资金池付款状态
func Tmallfanscashpoolcheckpay(clt *core.SDKClient, req *fans.TmallfanscashpoolcheckpayAPIRequest, session string) (*fans.TmallfanscashpoolcheckpayAPIResponse, error) {
	var resp fans.TmallfanscashpoolcheckpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
