package alihealthmdeer

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmdeer"
)

// AlibabaalihealthmdeersciencesynVideo 视频同步【保存/更新】
// alibaba.alihealth.mdeer.science.synVideo
//
// 视频同步【保存/更新】
func AlibabaalihealthmdeersciencesynVideo(clt *core.SDKClient, req *alihealthmdeer.AlibabaalihealthmdeersciencesynVideoAPIRequest, session string) (*alihealthmdeer.AlibabaalihealthmdeersciencesynVideoAPIResponse, error) {
	var resp alihealthmdeer.AlibabaalihealthmdeersciencesynVideoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
