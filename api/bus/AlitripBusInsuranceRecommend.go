package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// AlitripBusInsuranceRecommend 汽车票保险推荐接口
// alitrip.bus.insurance.recommend
//
// 汽车票保险推荐接口-供商家售卖飞猪保险使用
func AlitripBusInsuranceRecommend(clt *core.SDKClient, req *bus.AlitripBusInsuranceRecommendAPIRequest, session string) (*bus.AlitripBusInsuranceRecommendAPIResponse, error) {
	var resp bus.AlitripBusInsuranceRecommendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
