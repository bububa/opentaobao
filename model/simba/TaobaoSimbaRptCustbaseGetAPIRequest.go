package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarptcustbasegetAPIRequest 客户账户报表基础数据对象 API请求
// taobao.simba.rpt.custbase.get
//
// 客户账户报表基础数据对象
type TaobaosimbarptcustbasegetAPIRequest struct {
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

// NewTaobaosimbarptcustbasegetRequest 初始化TaobaosimbarptcustbasegetAPIRequest对象
func NewTaobaosimbarptcustbasegetRequest() *TaobaosimbarptcustbasegetAPIRequest {
	return &TaobaosimbarptcustbasegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbarptcustbasegetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.custbase.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbarptcustbasegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbarptcustbasegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubwayToken is SubwayToken Setter
// 权限验证信息
func (r *TaobaosimbarptcustbasegetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// GetSubwayToken SubwayToken Getter
func (r TaobaosimbarptcustbasegetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaosimbarptcustbasegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbarptcustbasegetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始日期，格式yyyy-mm-dd
func (r *TaobaosimbarptcustbasegetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbarptcustbasegetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束日期，格式yyyy-mm-dd
func (r *TaobaosimbarptcustbasegetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosimbarptcustbasegetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetSource is Source Setter
// 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaosimbarptcustbasegetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaosimbarptcustbasegetAPIRequest) GetSource() string {
	return r._source
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaosimbarptcustbasegetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbarptcustbasegetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaosimbarptcustbasegetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbarptcustbasegetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
