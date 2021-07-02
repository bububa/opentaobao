package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripTrainCitySuggest 火车票城市搜索
// alitrip.btrip.train.city.suggest
//
// 阿里商旅提供火车票搜索接口，方便OA厂商更精准的对接
func AlitripBtripTrainCitySuggest(clt *core.SDKClient, req *btrip.AlitripBtripTrainCitySuggestAPIRequest, session string) (*btrip.AlitripBtripTrainCitySuggestAPIResponse, error) {
	var resp btrip.AlitripBtripTrainCitySuggestAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
