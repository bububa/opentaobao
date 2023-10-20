package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeVideoSync 视频草稿信息同步
// alibaba.alihouse.newhome.video.sync
//
// 接收视频信息记录
func AlibabaAlihouseNewhomeVideoSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeVideoSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeVideoSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
