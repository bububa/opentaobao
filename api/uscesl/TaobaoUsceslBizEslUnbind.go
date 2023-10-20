package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizEslUnbind 电子价签解绑接口
// taobao.uscesl.biz.esl.unbind
//
// 电子价签解绑接口
func TaobaoUsceslBizEslUnbind(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizEslUnbindAPIRequest, session string) (*uscesl.TaobaoUsceslBizEslUnbindAPIResponse, error) {
	var resp uscesl.TaobaoUsceslBizEslUnbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
