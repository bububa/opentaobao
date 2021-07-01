package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQtReportsGetAPIRequest
批量查询质检报告 API请求
taobao.qt.reports.get

批量查询质检报告，目前只支持查询qtType=11（天猫真假鉴定）类型的报告 */
type TaobaoQtReportsGetAPIRequest struct {
	model.Params
	// 质检服务商名
	_spName string
	// 质检类型，目前只支持查询qt_type=11的类型
	_qtType int64
	// 收费项code
	_servcieItemCode string
	// 送检者昵称
	_nick string
	// 查询时间段的开始时间
	_startTime string
	// 查询时间段的结束时间
	_endTime string
}

// NewTaobaoQtReportsGetRequest 初始化TaobaoQtReportsGetAPIRequest对象
func NewTaobaoQtReportsGetRequest() *TaobaoQtReportsGetAPIRequest {
	return &TaobaoQtReportsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQtReportsGetAPIRequest) GetApiMethodName() string {
	return "taobao.qt.reports.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQtReportsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SpName Setter
// 质检服务商名
func (r *TaobaoQtReportsGetAPIRequest) SetSpName(_spName string) error {
	r._spName = _spName
	r.Set("sp_name", _spName)
	return nil
}

// Get SpName Getter
func (r TaobaoQtReportsGetAPIRequest) GetSpName() string {
	return r._spName
}

// Set is QtType Setter
// 质检类型，目前只支持查询qt_type=11的类型
func (r *TaobaoQtReportsGetAPIRequest) SetQtType(_qtType int64) error {
	r._qtType = _qtType
	r.Set("qt_type", _qtType)
	return nil
}

// Get QtType Getter
func (r TaobaoQtReportsGetAPIRequest) GetQtType() int64 {
	return r._qtType
}

// Set is ServcieItemCode Setter
// 收费项code
func (r *TaobaoQtReportsGetAPIRequest) SetServcieItemCode(_servcieItemCode string) error {
	r._servcieItemCode = _servcieItemCode
	r.Set("servcie_item_code", _servcieItemCode)
	return nil
}

// Get ServcieItemCode Getter
func (r TaobaoQtReportsGetAPIRequest) GetServcieItemCode() string {
	return r._servcieItemCode
}

// Set is Nick Setter
// 送检者昵称
func (r *TaobaoQtReportsGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoQtReportsGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is StartTime Setter
// 查询时间段的开始时间
func (r *TaobaoQtReportsGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoQtReportsGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is EndTime Setter
// 查询时间段的结束时间
func (r *TaobaoQtReportsGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoQtReportsGetAPIRequest) GetEndTime() string {
	return r._endTime
}
