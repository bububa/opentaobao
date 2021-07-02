package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopTobPackageQueryAPIRequest B2B包裹查询接口 API请求
// alibaba.ascp.uop.tob.package.query
//
// 供应链中台TOB包裹查询接口
type AlibabaAscpUopTobPackageQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_packageQueryRequest *Packagequeryrequest
}

// NewAlibabaAscpUopTobPackageQueryRequest 初始化AlibabaAscpUopTobPackageQueryAPIRequest对象
func NewAlibabaAscpUopTobPackageQueryRequest() *AlibabaAscpUopTobPackageQueryAPIRequest {
	return &AlibabaAscpUopTobPackageQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopTobPackageQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.tob.package.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopTobPackageQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPackageQueryRequest is PackageQueryRequest Setter
// 系统自动生成
func (r *AlibabaAscpUopTobPackageQueryAPIRequest) SetPackageQueryRequest(_packageQueryRequest *Packagequeryrequest) error {
	r._packageQueryRequest = _packageQueryRequest
	r.Set("package_query_request", _packageQueryRequest)
	return nil
}

// GetPackageQueryRequest PackageQueryRequest Getter
func (r AlibabaAscpUopTobPackageQueryAPIRequest) GetPackageQueryRequest() *Packagequeryrequest {
	return r._packageQueryRequest
}
