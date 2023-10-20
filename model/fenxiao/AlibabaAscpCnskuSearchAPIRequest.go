package fenxiao

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpCnskuSearchAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpCnskuSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.cnsku.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpCnskuSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpCnskuSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAscpCnskuSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpCnskuSearchRequest()
	},
}

// GetAlibabaAscpCnskuSearchRequest 从 sync.Pool 获取 AlibabaAscpCnskuSearchAPIRequest
func GetAlibabaAscpCnskuSearchAPIRequest() *AlibabaAscpCnskuSearchAPIRequest {
	return poolAlibabaAscpCnskuSearchAPIRequest.Get().(*AlibabaAscpCnskuSearchAPIRequest)
}

// ReleaseAlibabaAscpCnskuSearchAPIRequest 将 AlibabaAscpCnskuSearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpCnskuSearchAPIRequest(v *AlibabaAscpCnskuSearchAPIRequest) {
	v.Reset()
	poolAlibabaAscpCnskuSearchAPIRequest.Put(v)
}
