package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkchannelorderusercancel 用户发起售中取消
// alibaba.wdk.channel.order.usercancel
//
// 用户发起售中取消
func Alibabawdkchannelorderusercancel(clt *core.SDKClient, req *wdk.AlibabawdkchannelorderusercancelAPIRequest, session string) (*wdk.AlibabawdkchannelorderusercancelAPIResponse, error) {
	var resp wdk.AlibabawdkchannelorderusercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
