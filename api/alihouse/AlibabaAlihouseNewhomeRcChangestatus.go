package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeRcChangestatus 图文草稿状态更新
// alibaba.alihouse.newhome.rc.changestatus
//
// 图文草稿状态更新
func AlibabaAlihouseNewhomeRcChangestatus(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRcChangestatusAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeRcChangestatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
