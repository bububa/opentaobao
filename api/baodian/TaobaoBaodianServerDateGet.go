package baodian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baodian"
)

// Taobaobaodianserverdateget 服务器时间获取
// taobao.baodian.server.date.get
//
// 获取服务器时间
func Taobaobaodianserverdateget(clt *core.SDKClient, req *baodian.TaobaobaodianserverdategetAPIRequest, session string) (*baodian.TaobaobaodianserverdategetAPIResponse, error) {
	var resp baodian.TaobaobaodianserverdategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
