package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaousceslbizeslunbind 电子价签解绑接口
// taobao.uscesl.biz.esl.unbind
//
// 电子价签解绑接口
func Taobaousceslbizeslunbind(clt *core.SDKClient, req *uscesl.TaobaousceslbizeslunbindAPIRequest, session string) (*uscesl.TaobaousceslbizeslunbindAPIResponse, error) {
	var resp uscesl.TaobaousceslbizeslunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
