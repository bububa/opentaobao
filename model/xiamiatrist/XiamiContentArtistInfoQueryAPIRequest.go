package xiamiatrist

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiContentArtistInfoQueryAPIRequest
搜索艺人列表 API请求
xiami.content.artist.info.query

根据查询条件，搜索相关艺人列表 */
type XiamiContentArtistInfoQueryAPIRequest struct {
	model.Params
	// 性别：1男性 2女性 3乐队
	_gender int64
	// 语种：1华语 2日本 3韩国 4欧美 5其他
	_language int64
	// 流派: 1嘻哈(说唱),2流行，3摇滚，4布鲁斯，5爵士，6雷鬼，7世界音乐，8拉丁，9电子，10节奏布鲁斯，11实验，12轻音乐，13新世纪，14舞台 / 银幕 / 娱乐      * 15唱作人，16民谣，18金属，19中国特色，20乡村，21古典，22儿童，23有声书，24动漫，25朋克
	_genre int64
	// 分页信息
	_pageReq *PagingVo
}

// New
