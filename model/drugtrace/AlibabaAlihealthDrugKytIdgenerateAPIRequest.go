package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytIdgenerateAPIRequest
终端(医疗机构|零售药店)ID生成接口 API请求
alibaba.alihealth.drug.kyt.idgenerate

终端(医疗机构|零售药店)ID生成接口 */
type AlibabaAlihealthDrugKytIdgenerateAPIRequest struct {
	model.Params
	// 行政区（省市区）
	_regionCode string
	// 零售药店、医疗机构名称
	_terminalName string
}

// NewAlibabaAlihealthDrugKytIdgenerateRequest 初始化AlibabaAlihealthDrugKytIdgenerateAPIRequest对象
func NewAlibabaAlihealthDrugKytIdgenerateRequest() *AlibabaAlihealthDrugKytIdgenerateAPIRequest {
	return &AlibabaAlihealthDrugKytIdgenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytIdgenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.idgenerate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytIdgenerateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RegionCode Setter
// 行政区（省市区）
func (r *AlibabaAlihealthDrugKytIdgenerateAPIRequest) SetRegionCode(_regionCode string) error {
	r._regionCode = _regionCode
	r.Set("region_code", _regionCode)
	return nil
}

// Get RegionCode Getter
func (r AlibabaAlihealthDrugKytIdgenerateAPIRequest) GetRegionCode() string {
	return r._regionCode
}

// Set is TerminalName Setter
// 零售药店、医疗机构名称
func (r *AlibabaAlihealthDrugKytIdgenerateAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// Get TerminalName Getter
func (r AlibabaAlihealthDrugKytIdgenerateAPIRequest) GetTerminalName() string {
	return r._terminalName
}
