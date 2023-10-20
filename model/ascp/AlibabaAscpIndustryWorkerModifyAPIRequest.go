package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustryworkermodifyAPIRequest 送货入户并安装修改师傅信息 API请求
// alibaba.ascp.industry.worker.modify
//
// 送货入户并安装修改师傅信息
type AlibabaascpindustryworkermodifyAPIRequest struct {
	model.Params
	// 请求对象
	_omsServiceModifyWorkerParameter *OmsServiceModifyWorkerParameter
}

// NewAlibabaascpindustryworkermodifyRequest 初始化AlibabaascpindustryworkermodifyAPIRequest对象
func NewAlibabaascpindustryworkermodifyRequest() *AlibabaascpindustryworkermodifyAPIRequest {
	return &AlibabaascpindustryworkermodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpindustryworkermodifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.worker.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpindustryworkermodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpindustryworkermodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOmsServiceModifyWorkerParameter is OmsServiceModifyWorkerParameter Setter
// 请求对象
func (r *AlibabaascpindustryworkermodifyAPIRequest) SetOmsServiceModifyWorkerParameter(_omsServiceModifyWorkerParameter *OmsServiceModifyWorkerParameter) error {
	r._omsServiceModifyWorkerParameter = _omsServiceModifyWorkerParameter
	r.Set("oms_service_modify_worker_parameter", _omsServiceModifyWorkerParameter)
	return nil
}

// GetOmsServiceModifyWorkerParameter OmsServiceModifyWorkerParameter Getter
func (r AlibabaascpindustryworkermodifyAPIRequest) GetOmsServiceModifyWorkerParameter() *OmsServiceModifyWorkerParameter {
	return r._omsServiceModifyWorkerParameter
}
