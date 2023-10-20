package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// AlitripBusInsuranceRecommend 汽车票保险推荐接口
// alitrip.bus.insurance.recommend
//
// 汽车票保险推荐接口-供商家售卖飞猪保险使用
func AlitripBusInsuranceRecommend(clt *core.SDKClient, req *bus.AlitripBusInsuranceRecommendAPIRequest, resp *bus.AlitripBusInsuranceRecommendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
