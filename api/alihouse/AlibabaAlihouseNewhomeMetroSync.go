package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeMetroSync 地铁数据同步
// alibaba.alihouse.newhome.metro.sync
//
// 地铁数据同步
func AlibabaAlihouseNewhomeMetroSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeMetroSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeMetroSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
