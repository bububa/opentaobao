package xiamicontent

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentMusicInfoQuery 搜索音乐
// xiami.content.music.info.query
//
// (批量)获取歌曲信息
func XiamiContentMusicInfoQuery(ctx context.Context, clt *core.SDKClient, req *xiamicontent.XiamiContentMusicInfoQueryAPIRequest, resp *xiamicontent.XiamiContentMusicInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
