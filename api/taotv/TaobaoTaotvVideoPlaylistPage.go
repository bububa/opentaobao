package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// Taobaotaotvvideoplaylistpage 分页获取所有播单
// taobao.taotv.video.playlist.page
//
// 获取所有播单信息（分页）
func Taobaotaotvvideoplaylistpage(clt *core.SDKClient, req *taotv.TaobaotaotvvideoplaylistpageAPIRequest, session string) (*taotv.TaobaotaotvvideoplaylistpageAPIResponse, error) {
	var resp taotv.TaobaotaotvvideoplaylistpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
