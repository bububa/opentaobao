package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpcnskusearchAPIRequest 供应链中台货品搜索接口 API请求
// alibaba.ascp.cnsku.search
//
// 供应链中台货品搜索接口
type AlibabaascpcnskusearchAPIRequest struct {
	model.Params
	// 货品搜索条件
	_param0 *ScItemSearchDto
}

// NewAlibabaascpcnskusearchRequest 初始化AlibabaascpcnskusearchAPIRequest对象
func NewAlibabaascpcnskusearchRequest() *AlibabaascpcnskusearchAPIRequest {
	return &AlibabaascpcnskusearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpcnskusearchAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.cnsku.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpcnskusearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpcnskusearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 货品搜索条件
func (r *AlibabaascpcnskusearchAPIRequest) SetParam0(_param0 *ScItemSearchDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaascpcnskusearchAPIRequest) GetParam0() *ScItemSearchDto {
	return r._param0
}
