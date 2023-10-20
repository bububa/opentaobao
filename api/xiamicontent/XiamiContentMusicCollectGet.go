package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentMusicCollectGet 获取歌单歌曲
// xiami.content.music.collect.get
//
// 获取歌单歌曲
func XiamiContentMusicCollectGet(clt *core.SDKClient, req *xiamicontent.XiamiContentMusicCollectGetAPIRequest, resp *xiamicontent.XiamiContentMusicCollectGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
