package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptCustbaseGetAPIRequest 客户账户报表基础数据对象 API请求
// taobao.simba.rpt.custbase.get
//
// 客户账户报表基础数据对象
type TaobaoSimbaRptCustbaseGetAPIRequest struct {
	model.Params
	// 权限验证信息
	_subwayToken string
	// 昵称
	_nick string
	// 开始日期，格式yyyy-mm-dd
	_startTime string
	// 结束日期，格式yyyy-mm-dd
	_endTime string
	// 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
	_source string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
}

// NewTaobaoSimbaRptCustbaseGetRequest 初始化TaobaoSimbaRptCustbaseGetAPIRequest对象
func NewTaobaoSimbaRptCustbaseGetRequest() *TaobaoSimbaRptCustbaseGetAPIRequest {
	return &TaobaoSimbaRptCustbaseGetAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) Reset() {
	r._subwayToken = ""
	r._nick = ""
	r._startTime = ""
	r._endTime = ""
	r._source = ""
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.custbase.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubwayToken is SubwayToken Setter
// 权限验证信息
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// GetSubwayToken SubwayToken Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetSource is Source Setter
// 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetSource() string {
	return r._source
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoSimbaRptCustbaseGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaRptCustbaseGetRequest()
	},
}

// GetTaobaoSimbaRptCustbaseGetRequest 从 sync.Pool 获取 TaobaoSimbaRptCustbaseGetAPIRequest
func GetTaobaoSimbaRptCustbaseGetAPIRequest() *TaobaoSimbaRptCustbaseGetAPIRequest {
	return poolTaobaoSimbaRptCustbaseGetAPIRequest.Get().(*TaobaoSimbaRptCustbaseGetAPIRequest)
}

// ReleaseTaobaoSimbaRptCustbaseGetAPIRequest 将 TaobaoSimbaRptCustbaseGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaRptCustbaseGetAPIRequest(v *TaobaoSimbaRptCustbaseGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaRptCustbaseGetAPIRequest.Put(v)
}
