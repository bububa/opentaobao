package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentAlbumInfoGet 获取专辑信息
// xiami.content.album.info.get
//
// 获取专辑信息
func XiamiContentAlbumInfoGet(clt *core.SDKClient, req *xiamicontent.XiamiContentAlbumInfoGetAPIRequest, session string) (*xiamicontent.XiamiContentAlbumInfoGetAPIResponse, error) {
	var resp xiamicontent.XiamiContentAlbumInfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
