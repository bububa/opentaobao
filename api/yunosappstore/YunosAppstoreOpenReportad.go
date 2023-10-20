package yunosappstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosappstore"
)

// Yunosappstoreopenreportad 外投广告上报接口
// yunos.appstore.open.reportad
//
// 外投广告回流上报接口
func Yunosappstoreopenreportad(clt *core.SDKClient, req *yunosappstore.YunosappstoreopenreportadAPIRequest, session string) (*yunosappstore.YunosappstoreopenreportadAPIResponse, error) {
	var resp yunosappstore.YunosappstoreopenreportadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
