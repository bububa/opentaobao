package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticLabelPrintAPIRequest 物流打印面单 API请求
// aliexpress.local.logistic.label.print
//
// 物流打印面单
type AliexpressLocalLogisticLabelPrintAPIRequest struct {
	model.Params
	// 打印面单入参
	_param1 *PrintLabelRequestDto
}

// NewAliexpressLocalLogisticLabelPrintRequest 初始化AliexpressLocalLogisticLabelPrintAPIRequest对象
func NewAliexpressLocalLogisticLabelPrintRequest() *AliexpressLocalLogisticLabelPrintAPIRequest {
	return &AliexpressLocalLogisticLabelPrintAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressLocalLogisticLabelPrintAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistic.label.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressLocalLogisticLabelPrintAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressLocalLogisticLabelPrintAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 打印面单入参
func (r *AliexpressLocalLogisticLabelPrintAPIRequest) SetParam1(_param1 *PrintLabelRequestDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpressLocalLogisticLabelPrintAPIRequest) GetParam1() *PrintLabelRequestDto {
	return r._param1
}
