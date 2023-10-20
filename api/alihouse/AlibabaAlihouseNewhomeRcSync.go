package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeRcSync 阿里房产图文草稿信息同步
// alibaba.alihouse.newhome.rc.sync
//
// 接收图文草稿信息
func AlibabaAlihouseNewhomeRcSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRcSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeRcSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
