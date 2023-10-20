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
	_reportdate string
	// 人群标签id
	_ucrowdid int64
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

// SetReportdate is Reportdate Setter
// 报表日期，支持：月、日，示例：2020-12，2020-12-30
func (r *TaobaotbkscucrowdreportgetAPIRequest) SetReportdate(_reportdate string) error {
	r._reportdate = _reportdate
	r.Set("report_date", _reportdate)
	return nil
}

// GetReportdate Reportdate Getter
func (r TaobaotbkscucrowdreportgetAPIRequest) GetReportdate() string {
	return r._reportdate
}

// SetUcrowdid is Ucrowdid Setter
// 人群标签id
func (r *TaobaotbkscucrowdreportgetAPIRequest) SetUcrowdid(_ucrowdid int64) error {
	r._ucrowdid = _ucrowdid
	r.Set("ucrowd_id", _ucrowdid)
	return nil
}

// GetUcrowdid Ucrowdid Getter
func (r TaobaotbkscucrowdreportgetAPIRequest) GetUcrowdid() int64 {
	return r._ucrowdid
}
