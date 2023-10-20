package xiamiopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamiopen"
)

// XiamiApiSongDetailGet 获取歌曲详情
// xiami.api.song.detail.get
//
// 获取歌曲详情
func XiamiApiSongDetailGet(clt *core.SDKClient, req *xiamiopen.XiamiApiSongDetailGetAPIRequest, session string) (*xiamiopen.XiamiApiSongDetailGetAPIResponse, error) {
	var resp xiamiopen.XiamiApiSongDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
