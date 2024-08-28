package xiamicontent

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentAlbumInfoGet 获取专辑信息
// xiami.content.album.info.get
//
// 获取专辑信息
func XiamiContentAlbumInfoGet(ctx context.Context, clt *core.SDKClient, req *xiamicontent.XiamiContentAlbumInfoGetAPIRequest, resp *xiamicontent.XiamiContentAlbumInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
