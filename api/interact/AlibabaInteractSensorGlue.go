package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensorglue 视频播放
// alibaba.interact.sensor.glue
//
// 视频播放
func Alibabainteractsensorglue(clt *core.SDKClient, req *interact.AlibabainteractsensorglueAPIRequest, session string) (*interact.AlibabainteractsensorglueAPIResponse, error) {
	var resp interact.AlibabainteractsensorglueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
