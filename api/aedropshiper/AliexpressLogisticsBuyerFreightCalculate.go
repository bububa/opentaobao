package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpresslogisticsbuyerfreightcalculate 提供给买家使用的运费计算接口
// aliexpress.logistics.buyer.freight.calculate
//
// 提供给买家使用的运费计算接口
func Aliexpresslogisticsbuyerfreightcalculate(clt *core.SDKClient, req *aedropshiper.AliexpresslogisticsbuyerfreightcalculateAPIRequest, session string) (*aedropshiper.AliexpresslogisticsbuyerfreightcalculateAPIResponse, error) {
	var resp aedropshiper.AliexpresslogisticsbuyerfreightcalculateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
