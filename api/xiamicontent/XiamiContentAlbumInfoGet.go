package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentAlbumInfoGet 获取专辑信息
// xiami.content.album.info.get
//
// 获取专辑信息
func XiamiContentAlbumInfoGet(clt *core.SDKClient, req *xiamicontent.XiamiContentAlbumInfoGetAPIRequest, resp *xiamicontent.XiamiContentAlbumInfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
