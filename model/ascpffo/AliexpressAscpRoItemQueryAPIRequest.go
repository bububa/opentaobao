package ascpffo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpRoItemQueryAPIRequest AliExpress退供单明细查询API API请求
// aliexpress.ascp.ro.item.query
//
// AE仓发 单个退供单明细查询
type AliexpressAscpRoItemQueryAPIRequest struct {
	model.Params
	// dto
	_returnOrderItemQuery *ReturnOrderItemQueryDto
}

// NewAliexpressAscpRoItemQueryRequest 初始化AliexpressAscpRoItemQueryAPIRequest对象
func NewAliexpressAscpRoItemQueryRequest() *AliexpressAscpRoItemQueryAPIRequest {
	return &AliexpressAscpRoItemQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAscpRoItemQueryAPIRequest) Reset() {
	r._returnOrderItemQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpRoItemQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.ro.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpRoItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpRoItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReturnOrderItemQuery is ReturnOrderItemQuery Setter
// dto
func (r *AliexpressAscpRoItemQueryAPIRequest) SetReturnOrderItemQuery(_returnOrderItemQuery *ReturnOrderItemQueryDto) error {
	r._returnOrderItemQuery = _returnOrderItemQuery
	r.Set("return_order_item_query", _returnOrderItemQuery)
	return nil
}

// GetReturnOrderItemQuery ReturnOrderItemQuery Getter
func (r AliexpressAscpRoItemQueryAPIRequest) GetReturnOrderItemQuery() *ReturnOrderItemQueryDto {
	return r._returnOrderItemQuery
}

var poolAliexpressAscpRoItemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAscpRoItemQueryRequest()
	},
}

// GetAliexpressAscpRoItemQueryRequest 从 sync.Pool 获取 AliexpressAscpRoItemQueryAPIRequest
func GetAliexpressAscpRoItemQueryAPIRequest() *AliexpressAscpRoItemQueryAPIRequest {
	return poolAliexpressAscpRoItemQueryAPIRequest.Get().(*AliexpressAscpRoItemQueryAPIRequest)
}

// ReleaseAliexpressAscpRoItemQueryAPIRequest 将 AliexpressAscpRoItemQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressAscpRoItemQueryAPIRequest(v *AliexpressAscpRoItemQueryAPIRequest) {
	v.Reset()
	poolAliexpressAscpRoItemQueryAPIRequest.Put(v)
}
