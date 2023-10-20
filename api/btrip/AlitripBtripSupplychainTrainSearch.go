package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainTrainSearch 【商旅】火车票订单查询
// alitrip.btrip.supplychain.train.search
//
// 【商旅】火车票订单查询
func AlitripBtripSupplychainTrainSearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainTrainSearchAPIRequest, resp *btrip.AlitripBtripSupplychainTrainSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
