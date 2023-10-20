package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// Aliyunindustrytttmorderquery 天天特卖数字工厂订单获取
// aliyun.industry.tttm.order.query
//
// 获取阿里云数字工厂内天天特卖业务的订单
func Aliyunindustrytttmorderquery(clt *core.SDKClient, req *tttm.AliyunindustrytttmorderqueryAPIRequest, session string) (*tttm.AliyunindustrytttmorderqueryAPIResponse, error) {
	var resp tttm.AliyunindustrytttmorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
