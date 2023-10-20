package mc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mc"
)

// AliyunUnimktTaskChargeLaunch 云码权益查询
// aliyun.unimkt.task.charge.launch
//
// 云码线上流量投放链路，用于判断用户是否有匹配的投放计划
func AliyunUnimktTaskChargeLaunch(clt *core.SDKClient, req *mc.AliyunUnimktTaskChargeLaunchAPIRequest, resp *mc.AliyunUnimktTaskChargeLaunchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
