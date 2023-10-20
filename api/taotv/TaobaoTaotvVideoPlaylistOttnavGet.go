package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// Taobaotaotvvideoplaylistottnavget ott播单
// taobao.taotv.video.playlist.ottnav.get
//
// 根据聚焦播单ID拿到下面播单视频，根据左侧播单ID列表批量拿到播单信息
func Taobaotaotvvideoplaylistottnavget(clt *core.SDKClient, req *taotv.TaobaotaotvvideoplaylistottnavgetAPIRequest, session string) (*taotv.TaobaotaotvvideoplaylistottnavgetAPIResponse, error) {
	var resp taotv.TaobaotaotvvideoplaylistottnavgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
