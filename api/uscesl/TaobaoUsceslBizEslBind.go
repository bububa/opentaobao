package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizEslBind 电子价签绑定接口
// taobao.uscesl.biz.esl.bind
//
// 电子价签商品绑定接口
func TaobaoUsceslBizEslBind(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizEslBindAPIRequest, resp *uscesl.TaobaoUsceslBizEslBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
