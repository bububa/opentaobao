package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderGetAPIRequest 交易订单详情查询 API请求
// alibaba.wdk.order.get
//
// 五道口三江单据查询接口
type AlibabaWdkOrderGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_idListQueryReq *IdListQueryRequest
}

// NewAlibabaWdkOrderGetRequest 初始化AlibabaWdkOrderGetAPIRequest对象
func NewAlibabaWdkOrderGetRequest() *AlibabaWdkOrderGetAPIRequest {
	return &AlibabaWdkOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkOrderGetAPIRequest) Reset() {
	r._idListQueryReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdListQueryReq is IdListQueryReq Setter
// 系统自动生成
func (r *AlibabaWdkOrderGetAPIRequest) SetIdListQueryReq(_idListQueryReq *IdListQueryRequest) error {
	r._idListQueryReq = _idListQueryReq
	r.Set("id_list_query_req", _idListQueryReq)
	return nil
}

// GetIdListQueryReq IdListQueryReq Getter
func (r AlibabaWdkOrderGetAPIRequest) GetIdListQueryReq() *IdListQueryRequest {
	return r._idListQueryReq
}

var poolAlibabaWdkOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkOrderGetRequest()
	},
}

// GetAlibabaWdkOrderGetRequest 从 sync.Pool 获取 AlibabaWdkOrderGetAPIRequest
func GetAlibabaWdkOrderGetAPIRequest() *AlibabaWdkOrderGetAPIRequest {
	return poolAlibabaWdkOrderGetAPIRequest.Get().(*AlibabaWdkOrderGetAPIRequest)
}

// ReleaseAlibabaWdkOrderGetAPIRequest 将 AlibabaWdkOrderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkOrderGetAPIRequest(v *AlibabaWdkOrderGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkOrderGetAPIRequest.Put(v)
}
