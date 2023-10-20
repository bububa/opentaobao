package yunosminiapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosminiapp"
)

// Yunosminiappactivitycall 调用活动接口
// yunos.miniapp.activity.call
//
// 用于小程序调用活动接口
func Yunosminiappactivitycall(clt *core.SDKClient, req *yunosminiapp.YunosminiappactivitycallAPIRequest, session string) (*yunosminiapp.YunosminiappactivitycallAPIResponse, error) {
	var resp yunosminiapp.YunosminiappactivitycallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
