package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Alitripbusinsurancerecommend 汽车票保险推荐接口
// alitrip.bus.insurance.recommend
//
// 汽车票保险推荐接口-供商家售卖飞猪保险使用
func Alitripbusinsurancerecommend(clt *core.SDKClient, req *bus.AlitripbusinsurancerecommendAPIRequest, session string) (*bus.AlitripbusinsurancerecommendAPIResponse, error) {
	var resp bus.AlitripbusinsurancerecommendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
