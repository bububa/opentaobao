package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiContentSongsInfoQueryAPIRequest
搜索歌曲列表 API请求
xiami.content.songs.info.query

多维度查询歌曲列表 */
type XiamiContentSongsInfoQueryAPIRequest struct {
	model.Params
	// 搜索条件 key支持songName/singerName/copyrightStatus/publishStatus/keyword
	_searchTerms []SearchTermsDto
	// tag搜索条件，tag尽量不要超过50个
	_tagOptional *SongCatsSearchDto
	// 排序,默认按照最新排序 1最新 2本周最热 3本月最热
	_orderBy int64
	// 分页信息
	_page *PagingVo
}

// New
