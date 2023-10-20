package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanitemsubscribedailyleftquery 查询当天可添加的余量
// taobao.baichuan.item.subscribe.daily.left.query
//
// 查询当天可添加的余量
func Taobaobaichuanitemsubscribedailyleftquery(clt *core.SDKClient, req *baichuan.TaobaobaichuanitemsubscribedailyleftqueryAPIRequest, session string) (*baichuan.TaobaobaichuanitemsubscribedailyleftqueryAPIResponse, error) {
	var resp baichuan.TaobaobaichuanitemsubscribedailyleftqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
