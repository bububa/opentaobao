package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountAliyuncsComCreateAliyunAccount20130701 创建阿里云账号
// account.aliyuncs.com.CreateAliyunAccount.2013-07-01
//
// 根据给定的阿里云账号，密码以及手机号创建阿里云账号
func AccountAliyuncsComCreateAliyunAccount20130701(clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateAliyunAccount20130701APIRequest, session string) (*aliyun.AccountAliyuncsComCreateAliyunAccount20130701APIResponse, error) {
	var resp aliyun.AccountAliyuncsComCreateAliyunAccount20130701APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
