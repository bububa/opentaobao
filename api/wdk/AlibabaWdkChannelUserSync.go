package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkchannelusersync 会员同步
// alibaba.wdk.channel.user.sync
//
// 会员同步
func Alibabawdkchannelusersync(clt *core.SDKClient, req *wdk.AlibabawdkchannelusersyncAPIRequest, session string) (*wdk.AlibabawdkchannelusersyncAPIResponse, error) {
	var resp wdk.AlibabawdkchannelusersyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
