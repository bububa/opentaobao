package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// Xiamicontentsongsinfoquery 搜索歌曲列表
// xiami.content.songs.info.query
//
// 多维度查询歌曲列表
func Xiamicontentsongsinfoquery(clt *core.SDKClient, req *xiamicontent.XiamicontentsongsinfoqueryAPIRequest, session string) (*xiamicontent.XiamicontentsongsinfoqueryAPIResponse, error) {
	var resp xiamicontent.XiamicontentsongsinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
