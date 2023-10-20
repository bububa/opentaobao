package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// Xiamicontentsongsaudioget 获取歌曲音频
// xiami.content.songs.audio.get
//
// 获取歌曲音频
func Xiamicontentsongsaudioget(clt *core.SDKClient, req *xiamicontent.XiamicontentsongsaudiogetAPIRequest, session string) (*xiamicontent.XiamicontentsongsaudiogetAPIResponse, error) {
	var resp xiamicontent.XiamicontentsongsaudiogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
