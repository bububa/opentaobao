package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentSongsCollectGet 获取歌单详情接口
// xiami.content.songs.collect.get
//
// 根据歌单id，获取歌单详情
func XiamiContentSongsCollectGet(clt *core.SDKClient, req *xiamicontent.XiamiContentSongsCollectGetAPIRequest, resp *xiamicontent.XiamiContentSongsCollectGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
