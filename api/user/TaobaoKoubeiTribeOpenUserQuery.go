package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaokoubeitribeopenuserquery 获取用户openId
// taobao.koubei.tribe.open.user.query
//
// 口碑综合体通过手机号码获取加密后的用户openId
func Taobaokoubeitribeopenuserquery(clt *core.SDKClient, req *user.TaobaokoubeitribeopenuserqueryAPIRequest, session string) (*user.TaobaokoubeitribeopenuserqueryAPIResponse, error) {
	var resp user.TaobaokoubeitribeopenuserqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
