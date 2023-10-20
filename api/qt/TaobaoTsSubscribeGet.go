package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// Taobaotssubscribeget 淘宝服务订购关系查询
// taobao.ts.subscribe.get
//
// ts订购关系状态查询. 暂只支持1口价服务.
func Taobaotssubscribeget(clt *core.SDKClient, req *qt.TaobaotssubscribegetAPIRequest, session string) (*qt.TaobaotssubscribegetAPIResponse, error) {
	var resp qt.TaobaotssubscribegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
