package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountAliyuncsComCreateAliyunAccountForBid2013_07_01 为bid用户创建账号
// account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01
//
// 给指定的bid创建账号，同时账号打上owner bid的标记
func AccountAliyuncsComCreateAliyunAccountForBid2013_07_01(clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIRequest, session string) (*aliyun.AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIResponse, error) {
	var resp aliyun.AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
