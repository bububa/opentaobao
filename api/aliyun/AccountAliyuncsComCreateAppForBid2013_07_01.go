package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

/* AccountAliyuncsComCreateAppForBid2013_07_01
给当前渠道下的用户创建app
account.aliyuncs.com.CreateAppForBid.2013-07-01

给自己渠道下的用户创建app */
func AccountAliyuncsComCreateAppForBid2013_07_01(clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateAppForBid2013_07_01APIRequest, session string) (*aliyun.AccountAliyuncsComCreateAppForBid2013_07_01APIResponse, error) {
	var resp aliyun.AccountAliyuncsComCreateAppForBid2013_07_01APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
