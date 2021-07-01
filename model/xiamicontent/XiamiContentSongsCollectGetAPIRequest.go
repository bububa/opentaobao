package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiContentSongsCollectGetAPIRequest
获取歌单详情接口 API请求
xiami.content.songs.collect.get

根据歌单id，获取歌单详情 */
type XiamiContentSongsCollectGetAPIRequest struct {
	model.Params
	// 歌单id
	_collectId int64
	// 分页信息
	_page *PagingVo
}

// New
