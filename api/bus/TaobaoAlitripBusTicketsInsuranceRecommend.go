package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaoalitripbusticketsinsurancerecommend 汽车票保险推荐
// taobao.alitrip.bus.tickets.insurance.recommend
//
// 获取推荐保险内容
func Taobaoalitripbusticketsinsurancerecommend(clt *core.SDKClient, req *bus.TaobaoalitripbusticketsinsurancerecommendAPIRequest, session string) (*bus.TaobaoalitripbusticketsinsurancerecommendAPIResponse, error) {
	var resp bus.TaobaoalitripbusticketsinsurancerecommendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
