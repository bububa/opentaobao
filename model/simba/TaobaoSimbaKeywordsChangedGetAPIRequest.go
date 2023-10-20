package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordschangedgetAPIRequest 分页获取修改过的关键词ID、宝贝id、修改时间 API请求
// taobao.simba.keywords.changed.get
//
// 分页获取修改过的关键词ID、宝贝id、修改时间
type TaobaosimbakeywordschangedgetAPIRequest struct {
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

// NewTaobaosimbakeywordschangedgetRequest 初始化TaobaosimbakeywordschangedgetAPIRequest对象
func NewTaobaosimbakeywordschangedgetRequest() *TaobaosimbakeywordschangedgetAPIRequest {
	return &TaobaosimbakeywordschangedgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordschangedgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.changed.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordschangedgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordschangedgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbakeywordschangedgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbakeywordschangedgetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaosimbakeywordschangedgetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbakeywordschangedgetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetPageSize is PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaosimbakeywordschangedgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbakeywordschangedgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaosimbakeywordschangedgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbakeywordschangedgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
