package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeSupportSync 周边配套数据同步
// alibaba.alihouse.newhome.support.sync
//
// 周边配套数据同步
func AlibabaAlihouseNewhomeSupportSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeSupportSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeSupportSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
