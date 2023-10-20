package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentMusicInfoGet 获取音乐信息
// xiami.content.music.info.get
//
// (批量)获取歌曲信息
func XiamiContentMusicInfoGet(clt *core.SDKClient, req *xiamicontent.XiamiContentMusicInfoGetAPIRequest, session string) (*xiamicontent.XiamiContentMusicInfoGetAPIResponse, error) {
	var resp xiamicontent.XiamiContentMusicInfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
