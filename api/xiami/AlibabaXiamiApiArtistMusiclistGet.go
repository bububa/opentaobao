package xiami

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiami"
)

// AlibabaXiamiApiArtistMusiclistGet 热门艺人
// alibaba.xiami.api.artist.musiclist.get
//
// 热门艺人
func AlibabaXiamiApiArtistMusiclistGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiArtistMusiclistGetAPIRequest, session string) (*xiami.AlibabaXiamiApiArtistMusiclistGetAPIResponse, error) {
	var resp xiami.AlibabaXiamiApiArtistMusiclistGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
