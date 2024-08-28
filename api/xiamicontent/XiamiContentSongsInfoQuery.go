package xiamicontent

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentSongsInfoQuery 搜索歌曲列表
// xiami.content.songs.info.query
//
// 多维度查询歌曲列表
func XiamiContentSongsInfoQuery(ctx context.Context, clt *core.SDKClient, req *xiamicontent.XiamiContentSongsInfoQueryAPIRequest, resp *xiamicontent.XiamiContentSongsInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
