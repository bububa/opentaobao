package xiamiopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamiopen"
)

// XiamiApiSongListenfileGet 获取歌曲试听文件
// xiami.api.song.listenfile.get
//
// 获取歌曲试听文件
func XiamiApiSongListenfileGet(clt *core.SDKClient, req *xiamiopen.XiamiApiSongListenfileGetAPIRequest, resp *xiamiopen.XiamiApiSongListenfileGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
