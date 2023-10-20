package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountaliyuncscomdeleteAppForBid20130701 运营商删除用户的appkey
// account.aliyuncs.com.DeleteAppForBid.2013-07-01
//
// 删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。
func AccountaliyuncscomdeleteAppForBid20130701(clt *core.SDKClient, req *aliyun.AccountaliyuncscomdeleteAppForBid20130701APIRequest, session string) (*aliyun.AccountaliyuncscomdeleteAppForBid20130701APIResponse, error) {
	var resp aliyun.AccountaliyuncscomdeleteAppForBid20130701APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
