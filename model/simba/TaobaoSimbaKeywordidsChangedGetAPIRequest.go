package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordidsChangedGetAPIRequest 获取修改的词ID API请求
// taobao.simba.keywordids.changed.get
//
// 获取修改的词ID
type TaobaoSimbaKeywordidsChangedGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 得到此时间点之后的数据，不能大于一个月
	_startTime string
	// 返回的每页数据量大小,默认200最大1000
	_pageSize int64
	// 返回的第几页数据，默认为1
	_pageNo int64
}

// NewTaobaoSimbaKeywordidsChangedGetRequest 初始化TaobaoSimbaKeywordidsChangedGetAPIRequest对象
func NewTaobaoSimbaKeywordidsChangedGetRequest() *TaobaoSimbaKeywordidsChangedGetAPIRequest {
	return &TaobaoSimbaKeywordidsChangedGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordidsChangedGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywordids.changed.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordidsChangedGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordidsChangedGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaKeywordidsChangedGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaKeywordidsChangedGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoSimbaKeywordidsChangedGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaKeywordidsChangedGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoSimbaKeywordidsChangedGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaKeywordidsChangedGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoSimbaKeywordidsChangedGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
