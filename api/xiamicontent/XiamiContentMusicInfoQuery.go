package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// Xiamicontentmusicinfoquery 搜索音乐
// xiami.content.music.info.query
//
// (批量)获取歌曲信息
func Xiamicontentmusicinfoquery(clt *core.SDKClient, req *xiamicontent.XiamicontentmusicinfoqueryAPIRequest, session string) (*xiamicontent.XiamicontentmusicinfoqueryAPIResponse, error) {
	var resp xiamicontent.XiamicontentmusicinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
