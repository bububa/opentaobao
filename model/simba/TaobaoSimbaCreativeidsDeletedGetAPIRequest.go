package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativeidsDeletedGetAPIRequest 获取删除的创意ID API请求
// taobao.simba.creativeids.deleted.get
//
// 获取删除的创意ID
type TaobaoSimbaCreativeidsDeletedGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 得到这个时间点之后的数据，不能大于一个月
	_startTime string
	// 返回的每页数据量大小,默认200最大1000
	_pageSize int64
	// 返回的第几页数据，默认为1
	_pageNo int64
}

// NewTaobaoSimbaCreativeidsDeletedGetRequest 初始化TaobaoSimbaCreativeidsDeletedGetAPIRequest对象
func NewTaobaoSimbaCreativeidsDeletedGetRequest() *TaobaoSimbaCreativeidsDeletedGetAPIRequest {
	return &TaobaoSimbaCreativeidsDeletedGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativeidsDeletedGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.creativeids.deleted.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativeidsDeletedGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCreativeidsDeletedGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativeidsDeletedGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCreativeidsDeletedGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 得到这个时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaCreativeidsDeletedGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSimbaCreativeidsDeletedGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetPageSize is PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaCreativeidsDeletedGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaCreativeidsDeletedGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaCreativeidsDeletedGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaCreativeidsDeletedGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
