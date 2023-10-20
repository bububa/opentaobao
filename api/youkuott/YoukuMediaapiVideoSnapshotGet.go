package youkuott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/youkuott"
)

// YoukuMediaapiVideoSnapshotGet 根据视频ID查询视频缩微图
// youku.mediaapi.video.snapshot.get
//
// 根据视频ID查询视频缩微图
func YoukuMediaapiVideoSnapshotGet(clt *core.SDKClient, req *youkuott.YoukuMediaapiVideoSnapshotGetAPIRequest, resp *youkuott.YoukuMediaapiVideoSnapshotGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
