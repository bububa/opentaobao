package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxiprovideraccountbalance 供应商渠道余额
// alibaba.happytrip.taxi.provider.account.balance
//
// 查询不同供应商不同渠道账户余额
func Alibabahappytriptaxiprovideraccountbalance(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxiprovideraccountbalanceAPIRequest, session string) (*happytrip.AlibabahappytriptaxiprovideraccountbalanceAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxiprovideraccountbalanceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
