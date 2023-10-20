package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjmemberhasbind 喵街会员是否绑定
// alibaba.mj.member.hasbind
//
// 喵街检测用户是否为数字化会员
func Alibabamjmemberhasbind(clt *core.SDKClient, req *mos.AlibabamjmemberhasbindAPIRequest, session string) (*mos.AlibabamjmemberhasbindAPIResponse, error) {
	var resp mos.AlibabamjmemberhasbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
