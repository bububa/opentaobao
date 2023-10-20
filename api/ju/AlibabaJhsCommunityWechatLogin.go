package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityWechatLogin 聚划算用增淘外社群登录
// alibaba.jhs.community.wechat.login
//
// 聚划算用增淘外社群登录
func AlibabaJhsCommunityWechatLogin(clt *core.SDKClient, req *ju.AlibabaJhsCommunityWechatLoginAPIRequest, session string) (*ju.AlibabaJhsCommunityWechatLoginAPIResponse, error) {
	var resp ju.AlibabaJhsCommunityWechatLoginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
