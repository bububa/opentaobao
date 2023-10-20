package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkchannelcommentcreate 差评导入
// alibaba.wdk.channel.comment.create
//
// 差评导入
func Alibabawdkchannelcommentcreate(clt *core.SDKClient, req *wdk.AlibabawdkchannelcommentcreateAPIRequest, session string) (*wdk.AlibabawdkchannelcommentcreateAPIResponse, error) {
	var resp wdk.AlibabawdkchannelcommentcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
