package alime

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alime"
)

// Taobaoalimeusertokenget 获取用户免登录令牌
// taobao.alime.user.token.get
//
// 根据第三账号信息获取用户的免登录令牌
func Taobaoalimeusertokenget(clt *core.SDKClient, req *alime.TaobaoalimeusertokengetAPIRequest, session string) (*alime.TaobaoalimeusertokengetAPIResponse, error) {
	var resp alime.TaobaoalimeusertokengetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
