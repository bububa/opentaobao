package xiamiopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamiopen"
)

// Xiamiapisongdetailget 获取歌曲详情
// xiami.api.song.detail.get
//
// 获取歌曲详情
func Xiamiapisongdetailget(clt *core.SDKClient, req *xiamiopen.XiamiapisongdetailgetAPIRequest, session string) (*xiamiopen.XiamiapisongdetailgetAPIResponse, error) {
	var resp xiamiopen.XiamiapisongdetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
