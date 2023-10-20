package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// Xiamicontentsongscollectget 获取歌单详情接口
// xiami.content.songs.collect.get
//
// 根据歌单id，获取歌单详情
func Xiamicontentsongscollectget(clt *core.SDKClient, req *xiamicontent.XiamicontentsongscollectgetAPIRequest, session string) (*xiamicontent.XiamicontentsongscollectgetAPIResponse, error) {
	var resp xiamicontent.XiamicontentsongscollectgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
