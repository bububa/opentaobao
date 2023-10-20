package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountaliyuncscomcreateApp20130701 给指定用户创建appkey
// account.aliyuncs.com.CreateApp.2013-07-01
//
// 为某个用户创建appkey
func AccountaliyuncscomcreateApp20130701(clt *core.SDKClient, req *aliyun.AccountaliyuncscomcreateApp20130701APIRequest, session string) (*aliyun.AccountaliyuncscomcreateApp20130701APIResponse, error) {
	var resp aliyun.AccountaliyuncscomcreateApp20130701APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
