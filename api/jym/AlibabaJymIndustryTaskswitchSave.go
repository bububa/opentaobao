package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymIndustryTaskswitchSave 行业信息系统开关
// alibaba.jym.industry.taskswitch.save
//
// VMOS回调交易猫行业信息系统
func AlibabaJymIndustryTaskswitchSave(clt *core.SDKClient, req *jym.AlibabaJymIndustryTaskswitchSaveAPIRequest, session string) (*jym.AlibabaJymIndustryTaskswitchSaveAPIResponse, error) {
	var resp jym.AlibabaJymIndustryTaskswitchSaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
