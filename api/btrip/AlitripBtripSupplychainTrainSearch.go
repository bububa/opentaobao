package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripSupplychainTrainSearch
【商旅】火车票订单查询
alitrip.btrip.supplychain.train.search

【商旅】火车票订单查询 */
func AlitripBtripSupplychainTrainSearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainTrainSearchAPIRequest, session string) (*btrip.AlitripBtripSupplychainTrainSearchAPIResponse, error) {
	var resp btrip.AlitripBtripSupplychainTrainSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
