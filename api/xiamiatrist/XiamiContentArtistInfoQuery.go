package xiamiatrist

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamiatrist"
)

// XiamiContentArtistInfoQuery 搜索艺人列表
// xiami.content.artist.info.query
//
// 根据查询条件，搜索相关艺人列表
func XiamiContentArtistInfoQuery(ctx context.Context, clt *core.SDKClient, req *xiamiatrist.XiamiContentArtistInfoQueryAPIRequest, resp *xiamiatrist.XiamiContentArtistInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
