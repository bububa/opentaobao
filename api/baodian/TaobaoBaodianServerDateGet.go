package baodian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baodian"
)

// TaobaoBaodianServerDateGet 服务器时间获取
// taobao.baodian.server.date.get
//
// 获取服务器时间
func TaobaoBaodianServerDateGet(clt *core.SDKClient, req *baodian.TaobaoBaodianServerDateGetAPIRequest, resp *baodian.TaobaoBaodianServerDateGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
