package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtriptraincitysuggest 火车票城市搜索
// alitrip.btrip.train.city.suggest
//
// 阿里商旅提供火车票搜索接口，方便OA厂商更精准的对接
func Alitripbtriptraincitysuggest(clt *core.SDKClient, req *btrip.AlitripbtriptraincitysuggestAPIRequest, session string) (*btrip.AlitripbtriptraincitysuggestAPIResponse, error) {
	var resp btrip.AlitripbtriptraincitysuggestAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
