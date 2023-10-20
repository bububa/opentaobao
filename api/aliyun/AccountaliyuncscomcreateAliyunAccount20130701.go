package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountAliyuncsComCreateAliyunAccount20130701 创建阿里云账号
// account.aliyuncs.com.CreateAliyunAccount.2013-07-01
//
// 根据给定的阿里云账号，密码以及手机号创建阿里云账号
func AccountAliyuncsComCreateAliyunAccount20130701(clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateAliyunAccount20130701APIRequest, resp *aliyun.AccountAliyuncsComCreateAliyunAccount20130701APIResponse, session string) error {
	return clt.Post(req, resp, session)
}
