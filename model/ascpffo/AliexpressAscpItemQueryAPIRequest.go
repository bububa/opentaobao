package ascpffo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpItemQueryAPIRequest AliExpress货品查询查询API API请求
// aliexpress.ascp.item.query
//
// AE货品查询API
type AliexpressAscpItemQueryAPIRequest struct {
	model.Params
	// DTO
	_scItemQuery *ScItemQueryDto
}

// NewAliexpressAscpItemQueryRequest 初始化AliexpressAscpItemQueryAPIRequest对象
func NewAliexpressAscpItemQueryRequest() *AliexpressAscpItemQueryAPIRequest {
	return &AliexpressAscpItemQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAscpItemQueryAPIRequest) Reset() {
	r._scItemQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpItemQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScItemQuery is ScItemQuery Setter
// DTO
func (r *AliexpressAscpItemQueryAPIRequest) SetScItemQuery(_scItemQuery *ScItemQueryDto) error {
	r._scItemQuery = _scItemQuery
	r.Set("sc_item_query", _scItemQuery)
	return nil
}

// GetScItemQuery ScItemQuery Getter
func (r AliexpressAscpItemQueryAPIRequest) GetScItemQuery() *ScItemQueryDto {
	return r._scItemQuery
}

var poolAliexpressAscpItemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAscpItemQueryRequest()
	},
}

// GetAliexpressAscpItemQueryRequest 从 sync.Pool 获取 AliexpressAscpItemQueryAPIRequest
func GetAliexpressAscpItemQueryAPIRequest() *AliexpressAscpItemQueryAPIRequest {
	return poolAliexpressAscpItemQueryAPIRequest.Get().(*AliexpressAscpItemQueryAPIRequest)
}

// ReleaseAliexpressAscpItemQueryAPIRequest 将 AliexpressAscpItemQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressAscpItemQueryAPIRequest(v *AliexpressAscpItemQueryAPIRequest) {
	v.Reset()
	poolAliexpressAscpItemQueryAPIRequest.Put(v)
}
