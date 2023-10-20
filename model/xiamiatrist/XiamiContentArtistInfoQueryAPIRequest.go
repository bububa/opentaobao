package xiamiatrist

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentartistinfoqueryAPIRequest 搜索艺人列表 API请求
// xiami.content.artist.info.query
//
// 根据查询条件，搜索相关艺人列表
type XiamicontentartistinfoqueryAPIRequest struct {
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

// NewXiamicontentartistinfoqueryRequest 初始化XiamicontentartistinfoqueryAPIRequest对象
func NewXiamicontentartistinfoqueryRequest() *XiamicontentartistinfoqueryAPIRequest {
	return &XiamicontentartistinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamicontentartistinfoqueryAPIRequest) GetApiMethodName() string {
	return "xiami.content.artist.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamicontentartistinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamicontentartistinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGender is Gender Setter
// 性别：1男性 2女性 3乐队
func (r *XiamicontentartistinfoqueryAPIRequest) SetGender(_gender int64) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// GetGender Gender Getter
func (r XiamicontentartistinfoqueryAPIRequest) GetGender() int64 {
	return r._gender
}

// SetLanguage is Language Setter
// 语种：1华语 2日本 3韩国 4欧美 5其他
func (r *XiamicontentartistinfoqueryAPIRequest) SetLanguage(_language int64) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r XiamicontentartistinfoqueryAPIRequest) GetLanguage() int64 {
	return r._language
}

// SetGenre is Genre Setter
// 流派: 1嘻哈(说唱),2流行，3摇滚，4布鲁斯，5爵士，6雷鬼，7世界音乐，8拉丁，9电子，10节奏布鲁斯，11实验，12轻音乐，13新世纪，14舞台 / 银幕 / 娱乐      * 15唱作人，16民谣，18金属，19中国特色，20乡村，21古典，22儿童，23有声书，24动漫，25朋克
func (r *XiamicontentartistinfoqueryAPIRequest) SetGenre(_genre int64) error {
	r._genre = _genre
	r.Set("genre", _genre)
	return nil
}

// GetGenre Genre Getter
func (r XiamicontentartistinfoqueryAPIRequest) GetGenre() int64 {
	return r._genre
}

// SetPageReq is PageReq Setter
// 分页信息
func (r *XiamicontentartistinfoqueryAPIRequest) SetPageReq(_pageReq *PagingVo) error {
	r._pageReq = _pageReq
	r.Set("page_req", _pageReq)
	return nil
}

// GetPageReq PageReq Getter
func (r XiamicontentartistinfoqueryAPIRequest) GetPageReq() *PagingVo {
	return r._pageReq
}
