package xiamicontent

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentMusicCollectGet 获取歌单歌曲
// xiami.content.music.collect.get
//
// 获取歌单歌曲
func XiamiContentMusicCollectGet(ctx context.Context, clt *core.SDKClient, req *xiamicontent.XiamiContentMusicCollectGetAPIRequest, resp *xiamicontent.XiamiContentMusicCollectGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
