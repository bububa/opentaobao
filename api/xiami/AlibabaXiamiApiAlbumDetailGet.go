package xiami

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiami"
)

// AlibabaXiamiApiAlbumDetailGet 虾米音乐专辑详情接口
// alibaba.xiami.api.album.detail.get
//
// 虾米音乐专辑详情接口
func AlibabaXiamiApiAlbumDetailGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiAlbumDetailGetAPIRequest, session string) (*xiami.AlibabaXiamiApiAlbumDetailGetAPIResponse, error) {
	var resp xiami.AlibabaXiamiApiAlbumDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
