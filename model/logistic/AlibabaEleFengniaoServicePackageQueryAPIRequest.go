package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaoservicepackagequeryAPIRequest 预采购服务包查询接口 API请求
// alibaba.ele.fengniao.service.package.query
//
// 查询门店所在经纬度可用服务包的接口
type AlibabaelefengniaoservicepackagequeryAPIRequest struct {
	model.Params
	// 入参
	_param *Param
}

// NewAlibabaelefengniaoservicepackagequeryRequest 初始化AlibabaelefengniaoservicepackagequeryAPIRequest对象
func NewAlibabaelefengniaoservicepackagequeryRequest() *AlibabaelefengniaoservicepackagequeryAPIRequest {
	return &AlibabaelefengniaoservicepackagequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaelefengniaoservicepackagequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.service.package.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaelefengniaoservicepackagequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaelefengniaoservicepackagequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaelefengniaoservicepackagequeryAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaelefengniaoservicepackagequeryAPIRequest) GetParam() *Param {
	return r._param
}
