package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentSongsInfoGet 获取歌曲信息
// xiami.content.songs.info.get
//
// (批量)获取歌曲信息
func XiamiContentSongsInfoGet(clt *core.SDKClient, req *xiamicontent.XiamiContentSongsInfoGetAPIRequest, resp *xiamicontent.XiamiContentSongsInfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
