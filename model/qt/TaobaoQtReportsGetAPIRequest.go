package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqtreportsgetAPIRequest 批量查询质检报告 API请求
// taobao.qt.reports.get
//
// 批量查询质检报告，目前只支持查询qtType=11（天猫真假鉴定）类型的报告
type TaobaoqtreportsgetAPIRequest struct {
	model.Params
	// 质检服务商名
	_spName string
	// 收费项code
	_servcieItemCode string
	// 送检者昵称
	_nick string
	// 查询时间段的开始时间
	_startTime string
	// 查询时间段的结束时间
	_endTime string
	// 质检类型，目前只支持查询qt_type=11的类型
	_qtType int64
}

// NewTaobaoqtreportsgetRequest 初始化TaobaoqtreportsgetAPIRequest对象
func NewTaobaoqtreportsgetRequest() *TaobaoqtreportsgetAPIRequest {
	return &TaobaoqtreportsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqtreportsgetAPIRequest) GetApiMethodName() string {
	return "taobao.qt.reports.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqtreportsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqtreportsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpName is SpName Setter
// 质检服务商名
func (r *TaobaoqtreportsgetAPIRequest) SetSpName(_spName string) error {
	r._spName = _spName
	r.Set("sp_name", _spName)
	return nil
}

// GetSpName SpName Getter
func (r TaobaoqtreportsgetAPIRequest) GetSpName() string {
	return r._spName
}

// SetServcieItemCode is ServcieItemCode Setter
// 收费项code
func (r *TaobaoqtreportsgetAPIRequest) SetServcieItemCode(_servcieItemCode string) error {
	r._servcieItemCode = _servcieItemCode
	r.Set("servcie_item_code", _servcieItemCode)
	return nil
}

// GetServcieItemCode ServcieItemCode Getter
func (r TaobaoqtreportsgetAPIRequest) GetServcieItemCode() string {
	return r._servcieItemCode
}

// SetNick is Nick Setter
// 送检者昵称
func (r *TaobaoqtreportsgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoqtreportsgetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 查询时间段的开始时间
func (r *TaobaoqtreportsgetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoqtreportsgetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 查询时间段的结束时间
func (r *TaobaoqtreportsgetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoqtreportsgetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetQtType is QtType Setter
// 质检类型，目前只支持查询qt_type=11的类型
func (r *TaobaoqtreportsgetAPIRequest) SetQtType(_qtType int64) error {
	r._qtType = _qtType
	r.Set("qt_type", _qtType)
	return nil
}

// GetQtType QtType Getter
func (r TaobaoqtreportsgetAPIRequest) GetQtType() int64 {
	return r._qtType
}
