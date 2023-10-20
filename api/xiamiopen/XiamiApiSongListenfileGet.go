package xiamiopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamiopen"
)

// Xiamiapisonglistenfileget 获取歌曲试听文件
// xiami.api.song.listenfile.get
//
// 获取歌曲试听文件
func Xiamiapisonglistenfileget(clt *core.SDKClient, req *xiamiopen.XiamiapisonglistenfilegetAPIRequest, session string) (*xiamiopen.XiamiapisonglistenfilegetAPIResponse, error) {
	var resp xiamiopen.XiamiapisonglistenfilegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
