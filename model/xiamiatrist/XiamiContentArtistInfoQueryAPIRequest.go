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

// NewXiamiContentArtistInfoQueryRequest 初始化XiamiContentArtistInfoQueryAPIRequest对象
func NewXiamiContentArtistInfoQueryRequest() *XiamiContentArtistInfoQueryAPIRequest {
	return &XiamiContentArtistInfoQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentArtistInfoQueryAPIRequest) GetApiMethodName() string {
	return "xiami.content.artist.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentArtistInfoQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Gender Setter
// 性别：1男性 2女性 3乐队
func (r *XiamiContentArtistInfoQueryAPIRequest) SetGender(_gender int64) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// Get Gender Getter
func (r XiamiContentArtistInfoQueryAPIRequest) GetGender() int64 {
	return r._gender
}

// Set is Language Setter
// 语种：1华语 2日本 3韩国 4欧美 5其他
func (r *XiamiContentArtistInfoQueryAPIRequest) SetLanguage(_language int64) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// Get Language Getter
func (r XiamiContentArtistInfoQueryAPIRequest) GetLanguage() int64 {
	return r._language
}

// Set is Genre Setter
// 流派: 1嘻哈(说唱),2流行，3摇滚，4布鲁斯，5爵士，6雷鬼，7世界音乐，8拉丁，9电子，10节奏布鲁斯，11实验，12轻音乐，13新世纪，14舞台 / 银幕 / 娱乐      * 15唱作人，16民谣，18金属，19中国特色，20乡村，21古典，22儿童，23有声书，24动漫，25朋克
func (r *XiamiContentArtistInfoQueryAPIRequest) SetGenre(_genre int64) error {
	r._genre = _genre
	r.Set("genre", _genre)
	return nil
}

// Get Genre Getter
func (r XiamiContentArtistInfoQueryAPIRequest) GetGenre() int64 {
	return r._genre
}

// Set is PageReq Setter
// 分页信息
func (r *XiamiContentArtistInfoQueryAPIRequest) SetPageReq(_pageReq *PagingVo) error {
	r._pageReq = _pageReq
	r.Set("page_req", _pageReq)
	return nil
}

// Get PageReq Getter
func (r XiamiContentArtistInfoQueryAPIRequest) GetPageReq() *PagingVo {
	return r._pageReq
}
