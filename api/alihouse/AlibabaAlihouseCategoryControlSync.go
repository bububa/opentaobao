package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousecategorycontrolsync 类目权限上翻
// alibaba.alihouse.category.control.sync
//
// 类目权限上翻
func Alibabaalihousecategorycontrolsync(clt *core.SDKClient, req *alihouse.AlibabaalihousecategorycontrolsyncAPIRequest, session string) (*alihouse.AlibabaalihousecategorycontrolsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousecategorycontrolsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
