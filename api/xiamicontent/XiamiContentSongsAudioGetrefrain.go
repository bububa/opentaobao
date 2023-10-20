package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// Xiamicontentsongsaudiogetrefrain 获取副歌信息
// xiami.content.songs.audio.getrefrain
//
// 获取歌曲音频副歌
func Xiamicontentsongsaudiogetrefrain(clt *core.SDKClient, req *xiamicontent.XiamicontentsongsaudiogetrefrainAPIRequest, session string) (*xiamicontent.XiamicontentsongsaudiogetrefrainAPIResponse, error) {
	var resp xiamicontent.XiamicontentsongsaudiogetrefrainAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
