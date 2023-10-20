package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest 作业小票查询接口 API请求
// alibaba.wdk.fulfill.batch.query.by.batchids
//
// 根据节点等条件查询履约单小票信息
type AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest struct {
	model.Params
	// 批次号列表
	_batchIds []string
	// 作业节点类型： WAREHOUSE：仓  DELIVERY_DOCK：配送站 SHOP：经营店
	_nodeType string
	// warehouseCode, 出库仓，由基础店仓维护，盒马全域统一,
	_nodeCode string
}

// NewAlibabaWdkFulfillBatchQueryByBatchidsRequest 初始化AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest对象
func NewAlibabaWdkFulfillBatchQueryByBatchidsRequest() *AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest {
	return &AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) Reset() {
	r._batchIds = r._batchIds[:0]
	r._nodeType = ""
	r._nodeCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.batch.query.by.batchids"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchIds is BatchIds Setter
// 批次号列表
func (r *AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) SetBatchIds(_batchIds []string) error {
	r._batchIds = _batchIds
	r.Set("batch_ids", _batchIds)
	return nil
}

// GetBatchIds BatchIds Getter
func (r AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) GetBatchIds() []string {
	return r._batchIds
}

// SetNodeType is NodeType Setter
// 作业节点类型： WAREHOUSE：仓  DELIVERY_DOCK：配送站 SHOP：经营店
func (r *AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) SetNodeType(_nodeType string) error {
	r._nodeType = _nodeType
	r.Set("node_type", _nodeType)
	return nil
}

// GetNodeType NodeType Getter
func (r AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) GetNodeType() string {
	return r._nodeType
}

// SetNodeCode is NodeCode Setter
// warehouseCode, 出库仓，由基础店仓维护，盒马全域统一,
func (r *AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) SetNodeCode(_nodeCode string) error {
	r._nodeCode = _nodeCode
	r.Set("node_code", _nodeCode)
	return nil
}

// GetNodeCode NodeCode Getter
func (r AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) GetNodeCode() string {
	return r._nodeCode
}

var poolAlibabaWdkFulfillBatchQueryByBatchidsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkFulfillBatchQueryByBatchidsRequest()
	},
}

// GetAlibabaWdkFulfillBatchQueryByBatchidsRequest 从 sync.Pool 获取 AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest
func GetAlibabaWdkFulfillBatchQueryByBatchidsAPIRequest() *AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest {
	return poolAlibabaWdkFulfillBatchQueryByBatchidsAPIRequest.Get().(*AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest)
}

// ReleaseAlibabaWdkFulfillBatchQueryByBatchidsAPIRequest 将 AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkFulfillBatchQueryByBatchidsAPIRequest(v *AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) {
	v.Reset()
	poolAlibabaWdkFulfillBatchQueryByBatchidsAPIRequest.Put(v)
}
