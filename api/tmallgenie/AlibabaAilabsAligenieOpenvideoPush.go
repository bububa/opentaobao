package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabsaligenieopenvideopush 天猫精灵内容库视频分集数据推送接口
// alibaba.ailabs.aligenie.openvideo.push
//
// 天猫精灵内容库视频分集数据推送接口
func Alibabaailabsaligenieopenvideopush(clt *core.SDKClient, req *tmallgenie.AlibabaailabsaligenieopenvideopushAPIRequest, session string) (*tmallgenie.AlibabaailabsaligenieopenvideopushAPIResponse, error) {
	var resp tmallgenie.AlibabaailabsaligenieopenvideopushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
