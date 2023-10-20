package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// Taobaotoponcetokenget 网关一次性token获取
// taobao.top.once.token.get
//
// 网关一次性token获取，对接文档:
func Taobaotoponcetokenget(clt *core.SDKClient, req *tbtrade.TaobaotoponcetokengetAPIRequest, session string) (*tbtrade.TaobaotoponcetokengetAPIResponse, error) {
	var resp tbtrade.TaobaotoponcetokengetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
