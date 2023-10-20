package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomesyncbrokerisreceivecustomer 经纪人接待状态变更
// alibaba.alihouse.existinghome.sync.broker.is.receive.customer
//
// 经纪人接待状态变更
func Alibabaalihouseexistinghomesyncbrokerisreceivecustomer(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomesyncbrokerisreceivecustomerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
