package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativesChangedGetAPIRequest 分页获取修改过的广告创意ID和修改时间 API请求
// taobao.simba.creatives.changed.get
//
// 分页获取修改过的广告创意ID和修改时间
type TaobaoSimbaCreativesChangedGetAPIRequest struct {
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

// NewTaobaoSimbaCreativesChangedGetRequest 初始化TaobaoSimbaCreativesChangedGetAPIRequest对象
func NewTaobaoSimbaCreativesChangedGetRequest() *TaobaoSimbaCreativesChangedGetAPIRequest {
	return &TaobaoSimbaCreativesChangedGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCreativesChangedGetAPIRequest) Reset() {
	r._nick = ""
	r._startTime = ""
	r._pageSize = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.creatives.changed.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativesChangedGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaCreativesChangedGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetPageSize is PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaCreativesChangedGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaCreativesChangedGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolTaobaoSimbaCreativesChangedGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCreativesChangedGetRequest()
	},
}

// GetTaobaoSimbaCreativesChangedGetRequest 从 sync.Pool 获取 TaobaoSimbaCreativesChangedGetAPIRequest
func GetTaobaoSimbaCreativesChangedGetAPIRequest() *TaobaoSimbaCreativesChangedGetAPIRequest {
	return poolTaobaoSimbaCreativesChangedGetAPIRequest.Get().(*TaobaoSimbaCreativesChangedGetAPIRequest)
}

// ReleaseTaobaoSimbaCreativesChangedGetAPIRequest 将 TaobaoSimbaCreativesChangedGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCreativesChangedGetAPIRequest(v *TaobaoSimbaCreativesChangedGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCreativesChangedGetAPIRequest.Put(v)
}
