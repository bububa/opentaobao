package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslocallogisticlabelprintAPIRequest 物流打印面单 API请求
// aliexpress.local.logistic.label.print
//
// 物流打印面单
type AliexpresslocallogisticlabelprintAPIRequest struct {
	model.Params
	// 打印面单入参
	_param1 *PrintLabelRequestDto
}

// NewAliexpresslocallogisticlabelprintRequest 初始化AliexpresslocallogisticlabelprintAPIRequest对象
func NewAliexpresslocallogisticlabelprintRequest() *AliexpresslocallogisticlabelprintAPIRequest {
	return &AliexpresslocallogisticlabelprintAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresslocallogisticlabelprintAPIRequest) GetApiMethodName() string {
	return "aliexpress.local.logistic.label.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresslocallogisticlabelprintAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresslocallogisticlabelprintAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 打印面单入参
func (r *AliexpresslocallogisticlabelprintAPIRequest) SetParam1(_param1 *PrintLabelRequestDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AliexpresslocallogisticlabelprintAPIRequest) GetParam1() *PrintLabelRequestDto {
	return r._param1
}
