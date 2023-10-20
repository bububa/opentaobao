package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabaiworkmcmsgsendmobile 发送消息给手机用户
// alibaba.iwork.mc.msg.sendmobile
//
// 给手机用户发送对应操作结果的消息
func Alibabaiworkmcmsgsendmobile(clt *core.SDKClient, req *campus.AlibabaiworkmcmsgsendmobileAPIRequest, session string) (*campus.AlibabaiworkmcmsgsendmobileAPIResponse, error) {
	var resp campus.AlibabaiworkmcmsgsendmobileAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
