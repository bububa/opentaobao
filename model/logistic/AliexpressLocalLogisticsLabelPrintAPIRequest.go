package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticsLabelPrintAPIRequest print label API请求
// aliexpress.local.logistics.label.print
//
// print label
type AliexpressLocalLogisticsLabelPrintAPIRequest struct {
	model.Params
	// param
	_param1 *PrintLabelRequestDto
}

// NewAliexpressLocalLogisticsLabelPrintRequest 初始化AliexpressLocalLogisticsLabelPrintAPIRequest对象
func NewAliexpressLocalLogisticsLabelPrintRequest() *AliexpressLocalLogisticsLabelPrintAPIRequest {
	return &AliexpressLocalLogisticsLabelPrintAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressLocalLogisticsLabelPrintAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.label.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressLocalLogisticsLabelPrintAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam1 is Param1 Setter
// param
func (r *AliexpressLocalLogisticsLabelPrintAPIRequest) SetParam1(_param1 *PrintLabelRequestDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpressLocalLogisticsLabelPrintAPIRequest) GetParam1() *PrintLabelRequestDto {
	return r._param1
}
