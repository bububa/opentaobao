package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaokoubeitribeopenverifycodeapply 口碑综合体手机号获取验证码
// taobao.koubei.tribe.open.verify.code.apply
//
// 口碑综合体通过手机号获取验证码对外开放接口
func Taobaokoubeitribeopenverifycodeapply(clt *core.SDKClient, req *user.TaobaokoubeitribeopenverifycodeapplyAPIRequest, session string) (*user.TaobaokoubeitribeopenverifycodeapplyAPIResponse, error) {
	var resp user.TaobaokoubeitribeopenverifycodeapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
