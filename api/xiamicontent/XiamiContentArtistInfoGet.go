package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentArtistInfoGet 获取艺人信息
// xiami.content.artist.info.get
//
// (批量)获取艺人信息
func XiamiContentArtistInfoGet(clt *core.SDKClient, req *xiamicontent.XiamiContentArtistInfoGetAPIRequest, session string) (*xiamicontent.XiamiContentArtistInfoGetAPIResponse, error) {
	var resp xiamicontent.XiamiContentArtistInfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
