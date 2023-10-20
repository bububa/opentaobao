package youkuott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/youkuott"
)

// Youkumediaapivideosnapshotget 根据视频ID查询视频缩微图
// youku.mediaapi.video.snapshot.get
//
// 根据视频ID查询视频缩微图
func Youkumediaapivideosnapshotget(clt *core.SDKClient, req *youkuott.YoukumediaapivideosnapshotgetAPIRequest, session string) (*youkuott.YoukumediaapivideosnapshotgetAPIResponse, error) {
	var resp youkuott.YoukumediaapivideosnapshotgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
