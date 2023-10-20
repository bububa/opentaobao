package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabsaligenieopenvideoalbumscenepush 视频专辑场景接入接口
// alibaba.ailabs.aligenie.openvideoalbum.scenepush
//
// 视频专辑场景接入接口
func Alibabaailabsaligenieopenvideoalbumscenepush(clt *core.SDKClient, req *tmallgenie.AlibabaailabsaligenieopenvideoalbumscenepushAPIRequest, session string) (*tmallgenie.AlibabaailabsaligenieopenvideoalbumscenepushAPIResponse, error) {
	var resp tmallgenie.AlibabaailabsaligenieopenvideoalbumscenepushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
