package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// Xiamicontentalbuminfoget 获取专辑信息
// xiami.content.album.info.get
//
// 获取专辑信息
func Xiamicontentalbuminfoget(clt *core.SDKClient, req *xiamicontent.XiamicontentalbuminfogetAPIRequest, session string) (*xiamicontent.XiamicontentalbuminfogetAPIResponse, error) {
	var resp xiamicontent.XiamicontentalbuminfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
