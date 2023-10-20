package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// Alibabajymindustrytaskswitchsave 行业信息系统开关
// alibaba.jym.industry.taskswitch.save
//
// VMOS回调交易猫行业信息系统
func Alibabajymindustrytaskswitchsave(clt *core.SDKClient, req *jym.AlibabajymindustrytaskswitchsaveAPIRequest, session string) (*jym.AlibabajymindustrytaskswitchsaveAPIResponse, error) {
	var resp jym.AlibabajymindustrytaskswitchsaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
