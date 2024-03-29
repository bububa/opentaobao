package ascpchannel

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopTobPackageQueryAPIRequest) Reset() {
	r._packageQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopTobPackageQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.tob.package.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopTobPackageQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopTobPackageQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAscpUopTobPackageQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopTobPackageQueryRequest()
	},
}

// GetAlibabaAscpUopTobPackageQueryRequest 从 sync.Pool 获取 AlibabaAscpUopTobPackageQueryAPIRequest
func GetAlibabaAscpUopTobPackageQueryAPIRequest() *AlibabaAscpUopTobPackageQueryAPIRequest {
	return poolAlibabaAscpUopTobPackageQueryAPIRequest.Get().(*AlibabaAscpUopTobPackageQueryAPIRequest)
}

// ReleaseAlibabaAscpUopTobPackageQueryAPIRequest 将 AlibabaAscpUopTobPackageQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopTobPackageQueryAPIRequest(v *AlibabaAscpUopTobPackageQueryAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopTobPackageQueryAPIRequest.Put(v)
}
