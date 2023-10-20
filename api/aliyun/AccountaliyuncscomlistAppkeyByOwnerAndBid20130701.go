package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountaliyuncscomlistAppkeyByOwnerAndBid20130701 根据渠道id和状态列出appkey
// account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01
//
// 根据渠道id和状态列出appkey
func AccountaliyuncscomlistAppkeyByOwnerAndBid20130701(clt *core.SDKClient, req *aliyun.AccountaliyuncscomlistAppkeyByOwnerAndBid20130701APIRequest, session string) (*aliyun.AccountaliyuncscomlistAppkeyByOwnerAndBid20130701APIResponse, error) {
	var resp aliyun.AccountaliyuncscomlistAppkeyByOwnerAndBid20130701APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
