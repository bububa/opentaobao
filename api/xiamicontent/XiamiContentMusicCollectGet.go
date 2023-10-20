package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// Xiamicontentmusiccollectget 获取歌单歌曲
// xiami.content.music.collect.get
//
// 获取歌单歌曲
func Xiamicontentmusiccollectget(clt *core.SDKClient, req *xiamicontent.XiamicontentmusiccollectgetAPIRequest, session string) (*xiamicontent.XiamicontentmusiccollectgetAPIResponse, error) {
	var resp xiamicontent.XiamicontentmusiccollectgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
