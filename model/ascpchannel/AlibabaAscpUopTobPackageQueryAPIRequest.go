package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuoptobpackagequeryAPIRequest B2B包裹查询接口 API请求
// alibaba.ascp.uop.tob.package.query
//
// 供应链中台TOB包裹查询接口
type AlibabaascpuoptobpackagequeryAPIRequest struct {
	model.Params
	// 系统自动生成
	_packageQueryRequest *Packagequeryrequest
}

// NewAlibabaascpuoptobpackagequeryRequest 初始化AlibabaascpuoptobpackagequeryAPIRequest对象
func NewAlibabaascpuoptobpackagequeryRequest() *AlibabaascpuoptobpackagequeryAPIRequest {
	return &AlibabaascpuoptobpackagequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpuoptobpackagequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.tob.package.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpuoptobpackagequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpuoptobpackagequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPackageQueryRequest is PackageQueryRequest Setter
// 系统自动生成
func (r *AlibabaascpuoptobpackagequeryAPIRequest) SetPackageQueryRequest(_packageQueryRequest *Packagequeryrequest) error {
	r._packageQueryRequest = _packageQueryRequest
	r.Set("package_query_request", _packageQueryRequest)
	return nil
}

// GetPackageQueryRequest PackageQueryRequest Getter
func (r AlibabaascpuoptobpackagequeryAPIRequest) GetPackageQueryRequest() *Packagequeryrequest {
	return r._packageQueryRequest
}
