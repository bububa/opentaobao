package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytidgenerateAPIRequest 终端(医疗机构|零售药店)ID生成接口 API请求
// alibaba.alihealth.drug.kyt.idgenerate
//
// 终端(医疗机构|零售药店)ID生成接口
type AlibabaalihealthdrugkytidgenerateAPIRequest struct {
	model.Params
	// 行政区（省市区）
	_regionCode string
	// 零售药店、医疗机构名称
	_terminalName string
}

// NewAlibabaalihealthdrugkytidgenerateRequest 初始化AlibabaalihealthdrugkytidgenerateAPIRequest对象
func NewAlibabaalihealthdrugkytidgenerateRequest() *AlibabaalihealthdrugkytidgenerateAPIRequest {
	return &AlibabaalihealthdrugkytidgenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytidgenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.idgenerate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytidgenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytidgenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRegionCode is RegionCode Setter
// 行政区（省市区）
func (r *AlibabaalihealthdrugkytidgenerateAPIRequest) SetRegionCode(_regionCode string) error {
	r._regionCode = _regionCode
	r.Set("region_code", _regionCode)
	return nil
}

// GetRegionCode RegionCode Getter
func (r AlibabaalihealthdrugkytidgenerateAPIRequest) GetRegionCode() string {
	return r._regionCode
}

// SetTerminalName is TerminalName Setter
// 零售药店、医疗机构名称
func (r *AlibabaalihealthdrugkytidgenerateAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// GetTerminalName TerminalName Getter
func (r AlibabaalihealthdrugkytidgenerateAPIRequest) GetTerminalName() string {
	return r._terminalName
}
