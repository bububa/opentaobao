package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// Xiamicontentmusicinfoget 获取音乐信息
// xiami.content.music.info.get
//
// (批量)获取歌曲信息
func Xiamicontentmusicinfoget(clt *core.SDKClient, req *xiamicontent.XiamicontentmusicinfogetAPIRequest, session string) (*xiamicontent.XiamicontentmusicinfogetAPIResponse, error) {
	var resp xiamicontent.XiamicontentmusicinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
