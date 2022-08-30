package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpCnskuSearchAPIRequest 供应链中台货品搜索接口 API请求
// alibaba.ascp.cnsku.search
//
// 供应链中台货品搜索接口
type AlibabaAscpCnskuSearchAPIRequest struct {
	model.Params
	// 货品搜索条件
	_param0 *ScItemSearchDto
}

// NewAlibabaAscpCnskuSearchRequest 初始化AlibabaAscpCnskuSearchAPIRequest对象
func NewAlibabaAscpCnskuSearchRequest() *AlibabaAscpCnskuSearchAPIRequest {
	return &AlibabaAscpCnskuSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpCnskuSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.cnsku.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpCnskuSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 货品搜索条件
func (r *AlibabaAscpCnskuSearchAPIRequest) SetParam0(_param0 *ScItemSearchDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaAscpCnskuSearchAPIRequest) GetParam0() *ScItemSearchDto {
	return r._param0
}
