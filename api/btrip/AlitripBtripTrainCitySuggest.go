package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripTrainCitySuggest 火车票城市搜索
// alitrip.btrip.train.city.suggest
//
// 阿里商旅提供火车票搜索接口，方便OA厂商更精准的对接
func AlitripBtripTrainCitySuggest(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripTrainCitySuggestAPIRequest, resp *btrip.AlitripBtripTrainCitySuggestAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
