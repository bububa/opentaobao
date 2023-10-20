package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// Taobaojymmemberrealnameverifywithoutuid 用户实名认证
// taobao.jym.member.realname.verify.withoutuid
//
// 开放用户实名认证接口使用
func Taobaojymmemberrealnameverifywithoutuid(clt *core.SDKClient, req *jym.TaobaojymmemberrealnameverifywithoutuidAPIRequest, session string) (*jym.TaobaojymmemberrealnameverifywithoutuidAPIResponse, error) {
	var resp jym.TaobaojymmemberrealnameverifywithoutuidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
