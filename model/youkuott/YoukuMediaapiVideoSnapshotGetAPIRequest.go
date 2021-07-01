package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuMediaapiVideoSnapshotGetAPIRequest
根据视频ID查询视频缩微图 API请求
youku.mediaapi.video.snapshot.get

根据视频ID查询视频缩微图 */
type YoukuMediaapiVideoSnapshotGetAPIRequest struct {
	model.Params
	// 视频id
	_vid string
}

// New
