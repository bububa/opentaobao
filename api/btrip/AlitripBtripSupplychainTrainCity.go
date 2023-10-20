package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainTrainCity 火车站数据查询
// alitrip.btrip.supplychain.train.city
//
// 火车站数据查询
func AlitripBtripSupplychainTrainCity(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainTrainCityAPIRequest, session string) (*btrip.AlitripBtripSupplychainTrainCityAPIResponse, error) {
	var resp btrip.AlitripBtripSupplychainTrainCityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
