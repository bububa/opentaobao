package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentMusicInfoGet 获取音乐信息
// xiami.content.music.info.get
//
// (批量)获取歌曲信息
func XiamiContentMusicInfoGet(clt *core.SDKClient, req *xiamicontent.XiamiContentMusicInfoGetAPIRequest, resp *xiamicontent.XiamiContentMusicInfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
