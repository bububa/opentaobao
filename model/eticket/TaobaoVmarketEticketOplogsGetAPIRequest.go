package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketOplogsGetAPIRequest 电子凭证操作日志查询 API请求
// taobao.vmarket.eticket.oplogs.get
//
// 电子凭证核销日志查询
type TaobaoVmarketEticketOplogsGetAPIRequest struct {
	model.Params
	// 0:全部 1:核销 2:冲正
	_type int64
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 核销码
	_code string
	// 手机号后四位
	_mobile string
	// 当前页码
	_pageNo int64
	// 每页显示的记录数，最大为40，默认为40
	_pageSize int64
	// 排序方式
	_sort string
	// 核销身份
	_posid string
	// 码商ID
	_codemerchantId int64
}

// NewTaobaoVmarketEticketOplogsGetRequest 初始化TaobaoVmarketEticketOplogsGetAPIRequest对象
func NewTaobaoVmarketEticketOplogsGetRequest() *TaobaoVmarketEticketOplogsGetAPIRequest {
	return &TaobaoVmarketEticketOplogsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.oplogs.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetType is Type Setter
// 0:全部 1:核销 2:冲正
func (r *TaobaoVmarketEticketOplogsGetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetType() int64 {
	return r._type
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoVmarketEticketOplogsGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoVmarketEticketOplogsGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCode is Code Setter
// 核销码
func (r *TaobaoVmarketEticketOplogsGetAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetCode() string {
	return r._code
}

// SetMobile is Mobile Setter
// 手机号后四位
func (r *TaobaoVmarketEticketOplogsGetAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetMobile() string {
	return r._mobile
}

// SetPageNo is PageNo Setter
// 当前页码
func (r *TaobaoVmarketEticketOplogsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页显示的记录数，最大为40，默认为40
func (r *TaobaoVmarketEticketOplogsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetSort is Sort Setter
// 排序方式
func (r *TaobaoVmarketEticketOplogsGetAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetSort() string {
	return r._sort
}

// SetPosid is Posid Setter
// 核销身份
func (r *TaobaoVmarketEticketOplogsGetAPIRequest) SetPosid(_posid string) error {
	r._posid = _posid
	r.Set("posid", _posid)
	return nil
}

// GetPosid Posid Getter
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetPosid() string {
	return r._posid
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID
func (r *TaobaoVmarketEticketOplogsGetAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaoVmarketEticketOplogsGetAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}
