package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// AliyunIndustryTttmOrderQuery 天天特卖数字工厂订单获取
// aliyun.industry.tttm.order.query
//
// 获取阿里云数字工厂内天天特卖业务的订单
func AliyunIndustryTttmOrderQuery(clt *core.SDKClient, req *tttm.AliyunIndustryTttmOrderQueryAPIRequest, resp *tttm.AliyunIndustryTttmOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
