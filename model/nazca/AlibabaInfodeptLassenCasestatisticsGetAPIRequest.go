package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInfodeptLassenCasestatisticsGetAPIRequest 法庭提交和结案案件量接口 API请求
// alibaba.infodept.lassen.casestatistics.get
//
// 功能描述：云嘉为浙江省高院制作数据大屏，需展示网上法庭相关数据，我方为省高院提供浙江省内法院收案和结案的案件量，开放数据接口，供其调取这两组数据。
type AlibabaInfodeptLassenCasestatisticsGetAPIRequest struct {
	model.Params
	// 地区代码
	_areaCode string
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
}

// NewAlibabaInfodeptLassenCasestatisticsGetRequest 初始化AlibabaInfodeptLassenCasestatisticsGetAPIRequest对象
func NewAlibabaInfodeptLassenCasestatisticsGetRequest() *AlibabaInfodeptLassenCasestatisticsGetAPIRequest {
	return &AlibabaInfodeptLassenCasestatisticsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInfodeptLassenCasestatisticsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.infodept.lassen.casestatistics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInfodeptLassenCasestatisticsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInfodeptLassenCasestatisticsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAreaCode is AreaCode Setter
// 地区代码
func (r *AlibabaInfodeptLassenCasestatisticsGetAPIRequest) SetAreaCode(_areaCode string) error {
	r._areaCode = _areaCode
	r.Set("area_code", _areaCode)
	return nil
}

// GetAreaCode AreaCode Getter
func (r AlibabaInfodeptLassenCasestatisticsGetAPIRequest) GetAreaCode() string {
	return r._areaCode
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *AlibabaInfodeptLassenCasestatisticsGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabaInfodeptLassenCasestatisticsGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *AlibabaInfodeptLassenCasestatisticsGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaInfodeptLassenCasestatisticsGetAPIRequest) GetEndTime() string {
	return r._endTime
}
