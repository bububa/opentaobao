package simba

import (
	"net/url"

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
	// 返回的每页数据量大小,默认200最大1000
	_pageSize int64
	// 返回的第几页数据，默认为1
	_pageNo int64
	// 得到此时间点之后的数据，不能大于一个月
	_startTime string
}

// NewTaobaoSimbaCreativesChangedGetRequest 初始化TaobaoSimbaCreativesChangedGetAPIRequest对象
func NewTaobaoSimbaCreativesChangedGetRequest() *TaobaoSimbaCreativesChangedGetAPIRequest {
	return &TaobaoSimbaCreativesChangedGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.creatives.changed.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativesChangedGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaCreativesChangedGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaCreativesChangedGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaCreativesChangedGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoSimbaCreativesChangedGetAPIRequest) GetStartTime() string {
	return r._startTime
}
