package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

/* XiamiContentSongsInfoGet
获取歌曲信息
xiami.content.songs.info.get

(批量)获取歌曲信息 */
func XiamiContentSongsInfoGet(clt *core.SDKClient, req *xiamicontent.XiamiContentSongsInfoGetAPIRequest, session string) (*xiamicontent.XiamiContentSongsInfoGetAPIResponse, error) {
	var resp xiamicontent.XiamiContentSongsInfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
