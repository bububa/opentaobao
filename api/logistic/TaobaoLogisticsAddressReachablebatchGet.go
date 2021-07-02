package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsAddressReachablebatchGet 批量判定服务是否可达
// taobao.logistics.address.reachablebatch.get
//
// 批量判定服务是否可达
func TaobaoLogisticsAddressReachablebatchGet(clt *core.SDKClient, req *logistic.TaobaoLogisticsAddressReachablebatchGetAPIRequest, session string) (*logistic.TaobaoLogisticsAddressReachablebatchGetAPIResponse, error) {
	var resp logistic.TaobaoLogisticsAddressReachablebatchGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
