package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsOrderGetAPIRequest 查询店仓作业单据清单 （库存对账辅助）-回流单 API请求
// alibaba.wdk.ums.order.get
//
// 查询店仓作业单据清单 （库存对账辅助）-回流单
type AlibabaWdkUmsOrderGetAPIRequest struct {
	model.Params
	// 查询单据的dto
	_queryErpbillDto *QueryErpBillDto
}

// NewAlibabaWdkUmsOrderGetRequest 初始化AlibabaWdkUmsOrderGetAPIRequest对象
func NewAlibabaWdkUmsOrderGetRequest() *AlibabaWdkUmsOrderGetAPIRequest {
	return &AlibabaWdkUmsOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkUmsOrderGetAPIRequest) Reset() {
	r._queryErpbillDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkUmsOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryErpbillDto is QueryErpbillDto Setter
// 查询单据的dto
func (r *AlibabaWdkUmsOrderGetAPIRequest) SetQueryErpbillDto(_queryErpbillDto *QueryErpBillDto) error {
	r._queryErpbillDto = _queryErpbillDto
	r.Set("query_erpbill_dto", _queryErpbillDto)
	return nil
}

// GetQueryErpbillDto QueryErpbillDto Getter
func (r AlibabaWdkUmsOrderGetAPIRequest) GetQueryErpbillDto() *QueryErpBillDto {
	return r._queryErpbillDto
}

var poolAlibabaWdkUmsOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkUmsOrderGetRequest()
	},
}

// GetAlibabaWdkUmsOrderGetRequest 从 sync.Pool 获取 AlibabaWdkUmsOrderGetAPIRequest
func GetAlibabaWdkUmsOrderGetAPIRequest() *AlibabaWdkUmsOrderGetAPIRequest {
	return poolAlibabaWdkUmsOrderGetAPIRequest.Get().(*AlibabaWdkUmsOrderGetAPIRequest)
}

// ReleaseAlibabaWdkUmsOrderGetAPIRequest 将 AlibabaWdkUmsOrderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkUmsOrderGetAPIRequest(v *AlibabaWdkUmsOrderGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkUmsOrderGetAPIRequest.Put(v)
}
