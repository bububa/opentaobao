package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainTrainDetailSearchVtwo 【商旅】火车票订单详情查询
// alitrip.btrip.supplychain.train.detail.search.vtwo
//
// 【商旅】火车票订单详情查询
func AlitripBtripSupplychainTrainDetailSearchVtwo(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest, resp *btrip.AlitripBtripSupplychainTrainDetailSearchVtwoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
