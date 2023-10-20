package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainTrainDetailSearch 【商旅】火车票订单详情查询
// alitrip.btrip.supplychain.train.detail.search
//
// 【商旅】火车票订单详情查询
func AlitripBtripSupplychainTrainDetailSearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainTrainDetailSearchAPIRequest, session string) (*btrip.AlitripBtripSupplychainTrainDetailSearchAPIResponse, error) {
	var resp btrip.AlitripBtripSupplychainTrainDetailSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
