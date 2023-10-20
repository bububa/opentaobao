package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaodeliverytemplatesget 获取用户下所有模板
// taobao.delivery.templates.get
//
// 根据用户ID获取用户下所有模板
func Taobaodeliverytemplatesget(clt *core.SDKClient, req *tblogistics.TaobaodeliverytemplatesgetAPIRequest, session string) (*tblogistics.TaobaodeliverytemplatesgetAPIResponse, error) {
	var resp tblogistics.TaobaodeliverytemplatesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
