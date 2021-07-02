package baodian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baodian"
)

// TaobaoBaodianServerDateGet 服务器时间获取
// taobao.baodian.server.date.get
//
// 获取服务器时间
func TaobaoBaodianServerDateGet(clt *core.SDKClient, req *baodian.TaobaoBaodianServerDateGetAPIRequest, session string) (*baodian.TaobaoBaodianServerDateGetAPIResponse, error) {
	var resp baodian.TaobaoBaodianServerDateGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
