package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieOpenvideoalbumScenepush 视频专辑场景接入接口
// alibaba.ailabs.aligenie.openvideoalbum.scenepush
//
// 视频专辑场景接入接口
func AlibabaAilabsAligenieOpenvideoalbumScenepush(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest, session string) (*tmallgenie.AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
