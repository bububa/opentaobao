package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// Xiamicontentsongsinfoget 获取歌曲信息
// xiami.content.songs.info.get
//
// (批量)获取歌曲信息
func Xiamicontentsongsinfoget(clt *core.SDKClient, req *xiamicontent.XiamicontentsongsinfogetAPIRequest, session string) (*xiamicontent.XiamicontentsongsinfogetAPIResponse, error) {
	var resp xiamicontent.XiamicontentsongsinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
