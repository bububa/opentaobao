package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Alibabavideopublish 发布视频
// alibaba.video.publish
//
// 发布视频。
// 说明：发布视频5s后再查询视频信息，否则可能无法获取播放链接
func Alibabavideopublish(clt *core.SDKClient, req *media.AlibabavideopublishAPIRequest, session string) (*media.AlibabavideopublishAPIResponse, error) {
	var resp media.AlibabavideopublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
