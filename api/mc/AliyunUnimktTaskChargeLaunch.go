package mc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mc"
)

// Aliyununimkttaskchargelaunch 云码权益查询
// aliyun.unimkt.task.charge.launch
//
// 云码线上流量投放链路，用于判断用户是否有匹配的投放计划
func Aliyununimkttaskchargelaunch(clt *core.SDKClient, req *mc.AliyununimkttaskchargelaunchAPIRequest, session string) (*mc.AliyununimkttaskchargelaunchAPIResponse, error) {
	var resp mc.AliyununimkttaskchargelaunchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
