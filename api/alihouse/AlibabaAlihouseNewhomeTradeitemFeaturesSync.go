package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeTradeitemFeaturesSync 同步品活动标
// alibaba.alihouse.newhome.tradeitem.features.sync
//
// 同步品活动标
func AlibabaAlihouseNewhomeTradeitemFeaturesSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
