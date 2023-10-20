package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabsbotsskilslist 对外设备获取技能列表
// alibaba.ailabs.bots.skils.list
//
// 获取ai开放平台技能列表
func Alibabaailabsbotsskilslist(clt *core.SDKClient, req *tmallgenie.AlibabaailabsbotsskilslistAPIRequest, session string) (*tmallgenie.AlibabaailabsbotsskilslistAPIResponse, error) {
	var resp tmallgenie.AlibabaailabsbotsskilslistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
