package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeVideoChangestatus 视频草稿状态更新
// alibaba.alihouse.newhome.video.changestatus
//
// 视频草稿状态更新
func AlibabaAlihouseNewhomeVideoChangestatus(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeVideoChangestatusAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeVideoChangestatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
