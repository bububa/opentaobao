package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscucrowdreportgetAPIRequest 淘宝客-服务商-人群推广效果 API请求
// taobao.tbk.sc.ucrowd.report.get
//
// 服务商使用。支持淘宝客通过入参人群标签id，获得人群的推广和转化效果数据。
type TaobaotbkscucrowdreportgetAPIRequest struct {
	model.Params
	// 报表日期，支持：月、日，示例：2020-12，2020-12-30
	_reportDate string
	// 人群标签id
	_ucrowdId int64
}

// NewTaobaotbkscucrowdreportgetRequest 初始化TaobaotbkscucrowdreportgetAPIRequest对象
func NewTaobaotbkscucrowdreportgetRequest() *TaobaotbkscucrowdreportgetAPIRequest {
	return &TaobaotbkscucrowdreportgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscucrowdreportgetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.ucrowd.report.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscucrowdreportgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscucrowdreportgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportDate is ReportDate Setter
// 报表日期，支持：月、日，示例：2020-12，2020-12-30
func (r *TaobaotbkscucrowdreportgetAPIRequest) SetReportDate(_reportDate string) error {
	r._reportDate = _reportDate
	r.Set("report_date", _reportDate)
	return nil
}

// GetReportDate ReportDate Getter
func (r TaobaotbkscucrowdreportgetAPIRequest) GetReportDate() string {
	return r._reportDate
}

// SetUcrowdId is UcrowdId Setter
// 人群标签id
func (r *TaobaotbkscucrowdreportgetAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaotbkscucrowdreportgetAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}
