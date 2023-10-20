package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjmemberbindmember 绑定会员
// alibaba.mj.member.bindmember
//
// 用于绑定喵街数字化会员
func Alibabamjmemberbindmember(clt *core.SDKClient, req *mos.AlibabamjmemberbindmemberAPIRequest, session string) (*mos.AlibabamjmemberbindmemberAPIResponse, error) {
	var resp mos.AlibabamjmemberbindmemberAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
