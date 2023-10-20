package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslocallogisticslabelprintAPIRequest print label API请求
// aliexpress.local.logistics.label.print
//
// print label
type AliexpresslocallogisticslabelprintAPIRequest struct {
	model.Params
	// param
	_param1 *PrintLabelRequestDto
}

// NewAliexpresslocallogisticslabelprintRequest 初始化AliexpresslocallogisticslabelprintAPIRequest对象
func NewAliexpresslocallogisticslabelprintRequest() *AliexpresslocallogisticslabelprintAPIRequest {
	return &AliexpresslocallogisticslabelprintAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresslocallogisticslabelprintAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistics.label.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresslocallogisticslabelprintAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresslocallogisticslabelprintAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// param
func (r *AliexpresslocallogisticslabelprintAPIRequest) SetParam1(_param1 *PrintLabelRequestDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpresslocallogisticslabelprintAPIRequest) GetParam1() *PrintLabelRequestDto {
	return r._param1
}
