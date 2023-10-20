package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeLayoutSync 房通户型数据同步
// alibaba.alihouse.newhome.layout.sync
//
// 房通户型数据同步
func AlibabaAlihouseNewhomeLayoutSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeLayoutSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeLayoutSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
