package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordidschangedgetAPIRequest 获取修改的词ID API请求
// taobao.simba.keywordids.changed.get
//
// 获取修改的词ID
type TaobaosimbakeywordidschangedgetAPIRequest struct {
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

// NewTaobaosimbakeywordidschangedgetRequest 初始化TaobaosimbakeywordidschangedgetAPIRequest对象
func NewTaobaosimbakeywordidschangedgetRequest() *TaobaosimbakeywordidschangedgetAPIRequest {
	return &TaobaosimbakeywordidschangedgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordidschangedgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywordids.changed.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordidschangedgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordidschangedgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbakeywordidschangedgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbakeywordidschangedgetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaosimbakeywordidschangedgetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbakeywordidschangedgetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetPageSize is PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaosimbakeywordidschangedgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbakeywordidschangedgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaosimbakeywordidschangedgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbakeywordidschangedgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
