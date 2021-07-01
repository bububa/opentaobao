package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

/* AlibabaAilabsAligenieVideoalbumPush
天猫精灵内容库视频合辑数据推送接口
alibaba.ailabs.aligenie.videoalbum.push

三方内容合作厂商可将视频辑数据通过此接口推送至天猫精灵内容库接入中，供天猫精灵使用 */
func AlibabaAilabsAligenieVideoalbumPush(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieVideoalbumPushAPIRequest, session string) (*tmallgenie.AlibabaAilabsAligenieVideoalbumPushAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsAligenieVideoalbumPushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
