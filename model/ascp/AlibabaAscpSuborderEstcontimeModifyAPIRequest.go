package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpsuborderestcontimemodifyAPIRequest 向前修改发货时效 API请求
// alibaba.ascp.suborder.estcontime.modify
//
// 向前修改发货时效
type AlibabaascpsuborderestcontimemodifyAPIRequest struct {
	model.Params
	// 修改商品发货时效请求
	_modifyEstConTimeRequest *ModifyEstConTimeRequest
}

// NewAlibabaascpsuborderestcontimemodifyRequest 初始化AlibabaascpsuborderestcontimemodifyAPIRequest对象
func NewAlibabaascpsuborderestcontimemodifyRequest() *AlibabaascpsuborderestcontimemodifyAPIRequest {
	return &AlibabaascpsuborderestcontimemodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpsuborderestcontimemodifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.suborder.estcontime.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpsuborderestcontimemodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpsuborderestcontimemodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyEstConTimeRequest is ModifyEstConTimeRequest Setter
// 修改商品发货时效请求
func (r *AlibabaascpsuborderestcontimemodifyAPIRequest) SetModifyEstConTimeRequest(_modifyEstConTimeRequest *ModifyEstConTimeRequest) error {
	r._modifyEstConTimeRequest = _modifyEstConTimeRequest
	r.Set("modify_est_con_time_request", _modifyEstConTimeRequest)
	return nil
}

// GetModifyEstConTimeRequest ModifyEstConTimeRequest Getter
func (r AlibabaascpsuborderestcontimemodifyAPIRequest) GetModifyEstConTimeRequest() *ModifyEstConTimeRequest {
	return r._modifyEstConTimeRequest
}
