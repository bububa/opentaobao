package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaousceslbizeslbind 电子价签绑定接口
// taobao.uscesl.biz.esl.bind
//
// 电子价签商品绑定接口
func Taobaousceslbizeslbind(clt *core.SDKClient, req *uscesl.TaobaousceslbizeslbindAPIRequest, session string) (*uscesl.TaobaousceslbizeslbindAPIResponse, error) {
	var resp uscesl.TaobaousceslbizeslbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
