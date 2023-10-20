package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Alibabainteractmediaaudio 音频相关鉴权接口
// alibaba.interact.media.audio
//
// 新音频包的鉴权接口
func Alibabainteractmediaaudio(clt *core.SDKClient, req *mtopopen.AlibabainteractmediaaudioAPIRequest, session string) (*mtopopen.AlibabainteractmediaaudioAPIResponse, error) {
	var resp mtopopen.AlibabainteractmediaaudioAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
