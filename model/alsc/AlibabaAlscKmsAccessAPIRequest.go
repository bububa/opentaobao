package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsckmsaccessAPIRequest 本地生活风控数据接入 API请求
// alibaba.alsc.kms.access
//
// 第三方使用本地生活数据对外提供服务，上报访问日志信息接口
type AlibabaalsckmsaccessAPIRequest struct {
	model.Params
	// 接入参数
	_requestdata string
}

// NewAlibabaalsckmsaccessRequest 初始化AlibabaalsckmsaccessAPIRequest对象
func NewAlibabaalsckmsaccessRequest() *AlibabaalsckmsaccessAPIRequest {
	return &AlibabaalsckmsaccessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsckmsaccessAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.kms.access"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsckmsaccessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsckmsaccessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestdata is Requestdata Setter
// 接入参数
func (r *AlibabaalsckmsaccessAPIRequest) SetRequestdata(_requestdata string) error {
	r._requestdata = _requestdata
	r.Set("requestdata", _requestdata)
	return nil
}

// GetRequestdata Requestdata Getter
func (r AlibabaalsckmsaccessAPIRequest) GetRequestdata() string {
	return r._requestdata
}
