package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiTagGenreSongGetAPIRequest 虾米-风格，流派歌曲 API请求
// alibaba.xiami.api.tag.genre.song.get
//
// 虾米-风格，流派歌曲
type AlibabaXiamiApiTagGenreSongGetAPIRequest struct {
	model.Params
	// 1:风格，2：流派
	_type string
	// 风格或流派id
	_id int64
	// 页数
	_page int64
	// 每页数量
	_limit int64
}

// NewAlibabaXiamiApiTagGenreSongGetRequest 初始化AlibabaXiamiApiTagGenreSongGetAPIRequest对象
func NewAlibabaXiamiApiTagGenreSongGetRequest() *AlibabaXiamiApiTagGenreSongGetAPIRequest {
	return &AlibabaXiamiApiTagGenreSongGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiTagGenreSongGetAPIRequest) GetApiMethodName() string {
	return "alibaba.xiami.api.tag.genre.song.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiTagGenreSongGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Type Setter
// 1:风格，2：流派
func (r *AlibabaXiamiApiTagGenreSongGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaXiamiApiTagGenreSongGetAPIRequest) GetType() string {
	return r._type
}

// Set is Id Setter
// 风格或流派id
func (r *AlibabaXiamiApiTagGenreSongGetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r AlibabaXiamiApiTagGenreSongGetAPIRequest) GetId() int64 {
	return r._id
}

// Set is Page Setter
// 页数
func (r *AlibabaXiamiApiTagGenreSongGetAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaXiamiApiTagGenreSongGetAPIRequest) GetPage() int64 {
	return r._page
}

// Set is Limit Setter
// 每页数量
func (r *AlibabaXiamiApiTagGenreSongGetAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// Get Limit Getter
func (r AlibabaXiamiApiTagGenreSongGetAPIRequest) GetLimit() int64 {
	return r._limit
}
