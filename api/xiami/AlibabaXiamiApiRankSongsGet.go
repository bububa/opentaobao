package xiami

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiami"
)

// AlibabaXiamiApiRankSongsGet 排行榜歌曲获取
// alibaba.xiami.api.rank.songs.get
//
// 获取歌曲排行榜
func AlibabaXiamiApiRankSongsGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiRankSongsGetAPIRequest, session string) (*xiami.AlibabaXiamiApiRankSongsGetAPIResponse, error) {
	var resp xiami.AlibabaXiamiApiRankSongsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
