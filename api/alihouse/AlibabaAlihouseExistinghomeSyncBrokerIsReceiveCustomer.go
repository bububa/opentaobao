package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomer 经纪人接待状态变更
// alibaba.alihouse.existinghome.sync.broker.is.receive.customer
//
// 经纪人接待状态变更
func AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomer(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
