package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoServicePackageQueryAPIRequest 预采购服务包查询接口 API请求
// alibaba.ele.fengniao.service.package.query
//
// 查询门店所在经纬度可用服务包的接口
type AlibabaEleFengniaoServicePackageQueryAPIRequest struct {
	model.Params
	// 入参
	_param *Param
}

// NewAlibabaEleFengniaoServicePackageQueryRequest 初始化AlibabaEleFengniaoServicePackageQueryAPIRequest对象
func NewAlibabaEleFengniaoServicePackageQueryRequest() *AlibabaEleFengniaoServicePackageQueryAPIRequest {
	return &AlibabaEleFengniaoServicePackageQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleFengniaoServicePackageQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoServicePackageQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.service.package.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoServicePackageQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoServicePackageQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaEleFengniaoServicePackageQueryAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaEleFengniaoServicePackageQueryAPIRequest) GetParam() *Param {
	return r._param
}

var poolAlibabaEleFengniaoServicePackageQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleFengniaoServicePackageQueryRequest()
	},
}

// GetAlibabaEleFengniaoServicePackageQueryRequest 从 sync.Pool 获取 AlibabaEleFengniaoServicePackageQueryAPIRequest
func GetAlibabaEleFengniaoServicePackageQueryAPIRequest() *AlibabaEleFengniaoServicePackageQueryAPIRequest {
	return poolAlibabaEleFengniaoServicePackageQueryAPIRequest.Get().(*AlibabaEleFengniaoServicePackageQueryAPIRequest)
}

// ReleaseAlibabaEleFengniaoServicePackageQueryAPIRequest 将 AlibabaEleFengniaoServicePackageQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleFengniaoServicePackageQueryAPIRequest(v *AlibabaEleFengniaoServicePackageQueryAPIRequest) {
	v.Reset()
	poolAlibabaEleFengniaoServicePackageQueryAPIRequest.Put(v)
}
