package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeVideoSync 视频草稿信息同步
// alibaba.alihouse.newhome.video.sync
//
// 接收视频信息记录
func AlibabaAlihouseNewhomeVideoSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeVideoSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeVideoSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeVideoSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
