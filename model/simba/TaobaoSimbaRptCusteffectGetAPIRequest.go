package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarptcusteffectgetAPIRequest 用户账户报表效果数据查询（只有汇总数据，无分类数据） API请求
// taobao.simba.rpt.custeffect.get
//
// 用户账户报表效果数据查询（只有汇总数据，无分类数据）
type TaobaosimbarptcusteffectgetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 开始时间，格式yyyy-mm-dd
	_startTime string
	// 结束时间，格式yyyy-mm-dd
	_endTime string
	// 权限校验参数
	_subwayToken string
	// 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
	_source string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
}

// NewTaobaosimbarptcusteffectgetRequest 初始化TaobaosimbarptcusteffectgetAPIRequest对象
func NewTaobaosimbarptcusteffectgetRequest() *TaobaosimbarptcusteffectgetAPIRequest {
	return &TaobaosimbarptcusteffectgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbarptcusteffectgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.custeffect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbarptcusteffectgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbarptcusteffectgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbarptcusteffectgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbarptcusteffectgetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaosimbarptcusteffectgetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbarptcusteffectgetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaosimbarptcusteffectgetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosimbarptcusteffectgetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetSubwayToken is SubwayToken Setter
// 权限校验参数
func (r *TaobaosimbarptcusteffectgetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// GetSubwayToken SubwayToken Getter
func (r TaobaosimbarptcusteffectgetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// SetSource is Source Setter
// 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaosimbarptcusteffectgetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaosimbarptcusteffectgetAPIRequest) GetSource() string {
	return r._source
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaosimbarptcusteffectgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbarptcusteffectgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaosimbarptcusteffectgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbarptcusteffectgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
