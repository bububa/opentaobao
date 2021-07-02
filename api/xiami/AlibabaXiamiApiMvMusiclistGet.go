package xiami

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiami"
)

// AlibabaXiamiApiMvMusiclistGet 乐馆mv列表
// alibaba.xiami.api.mv.musiclist.get
//
// 乐馆mv列表
func AlibabaXiamiApiMvMusiclistGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiMvMusiclistGetAPIRequest, session string) (*xiami.AlibabaXiamiApiMvMusiclistGetAPIResponse, error) {
	var resp xiami.AlibabaXiamiApiMvMusiclistGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
