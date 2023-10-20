package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeTradeitemFeaturesSync 同步品活动标
// alibaba.alihouse.newhome.tradeitem.features.sync
//
// 同步品活动标
func AlibabaAlihouseNewhomeTradeitemFeaturesSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
