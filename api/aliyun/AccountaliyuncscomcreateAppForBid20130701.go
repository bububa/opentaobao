package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountaliyuncscomcreateAppForBid20130701 给当前渠道下的用户创建app
// account.aliyuncs.com.CreateAppForBid.2013-07-01
//
// 给自己渠道下的用户创建app
func AccountaliyuncscomcreateAppForBid20130701(clt *core.SDKClient, req *aliyun.AccountaliyuncscomcreateAppForBid20130701APIRequest, session string) (*aliyun.AccountaliyuncscomcreateAppForBid20130701APIResponse, error) {
	var resp aliyun.AccountaliyuncscomcreateAppForBid20130701APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
