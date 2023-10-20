package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaxianyutenderorderperform 闲鱼暗拍订单履约
// alibaba.xianyu.tender.order.perform
//
// 闲鱼暗拍订单履约
func Alibabaxianyutenderorderperform(clt *core.SDKClient, req *idle.AlibabaxianyutenderorderperformAPIRequest, session string) (*idle.AlibabaxianyutenderorderperformAPIResponse, error) {
	var resp idle.AlibabaxianyutenderorderperformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
