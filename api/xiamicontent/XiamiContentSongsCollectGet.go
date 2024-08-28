package xiamicontent

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentSongsCollectGet 获取歌单详情接口
// xiami.content.songs.collect.get
//
// 根据歌单id，获取歌单详情
func XiamiContentSongsCollectGet(ctx context.Context, clt *core.SDKClient, req *xiamicontent.XiamiContentSongsCollectGetAPIRequest, resp *xiamicontent.XiamiContentSongsCollectGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
