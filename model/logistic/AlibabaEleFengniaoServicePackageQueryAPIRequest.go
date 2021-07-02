package logistic

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoServicePackageQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.service.package.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoServicePackageQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 入参
func (r *AlibabaEleFengniaoServicePackageQueryAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaEleFengniaoServicePackageQueryAPIRequest) GetParam() *Param {
	return r._param
}
