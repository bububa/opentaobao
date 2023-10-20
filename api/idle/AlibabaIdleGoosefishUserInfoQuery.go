package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlegoosefishuserinfoquery 闲鱼三方容器用户信息获取
// alibaba.idle.goosefish.user.info.query
//
// 闲鱼三方容器用户信息获取
func Alibabaidlegoosefishuserinfoquery(clt *core.SDKClient, req *idle.AlibabaidlegoosefishuserinfoqueryAPIRequest, session string) (*idle.AlibabaidlegoosefishuserinfoqueryAPIResponse, error) {
	var resp idle.AlibabaidlegoosefishuserinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
