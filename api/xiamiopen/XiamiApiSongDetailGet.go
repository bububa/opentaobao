package xiamiopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamiopen"
)

// XiamiApiSongDetailGet 获取歌曲详情
// xiami.api.song.detail.get
//
// 获取歌曲详情
func XiamiApiSongDetailGet(ctx context.Context, clt *core.SDKClient, req *xiamiopen.XiamiApiSongDetailGetAPIRequest, resp *xiamiopen.XiamiApiSongDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
