package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwayvideostateget 获取视频状态
// taobao.subway.video.state.get
//
// 获取已上传视频的状态
func Taobaosubwayvideostateget(clt *core.SDKClient, req *simba.TaobaosubwayvideostategetAPIRequest, session string) (*simba.TaobaosubwayvideostategetAPIResponse, error) {
	var resp simba.TaobaosubwayvideostategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
