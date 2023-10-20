package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfillbatchquerybybatchidsAPIRequest 作业小票查询接口 API请求
// alibaba.wdk.fulfill.batch.query.by.batchids
//
// 根据节点等条件查询履约单小票信息
type AlibabawdkfulfillbatchquerybybatchidsAPIRequest struct {
	model.Params
	// 批次号列表
	_batchIds []string
	// 作业节点类型： WAREHOUSE：仓  DELIVERY_DOCK：配送站 SHOP：经营店
	_nodeType string
	// warehouseCode, 出库仓，由基础店仓维护，盒马全域统一,
	_nodeCode string
}

// NewAlibabawdkfulfillbatchquerybybatchidsRequest 初始化AlibabawdkfulfillbatchquerybybatchidsAPIRequest对象
func NewAlibabawdkfulfillbatchquerybybatchidsRequest() *AlibabawdkfulfillbatchquerybybatchidsAPIRequest {
	return &AlibabawdkfulfillbatchquerybybatchidsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfillbatchquerybybatchidsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.batch.query.by.batchids"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfillbatchquerybybatchidsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfillbatchquerybybatchidsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchIds is BatchIds Setter
// 批次号列表
func (r *AlibabawdkfulfillbatchquerybybatchidsAPIRequest) SetBatchIds(_batchIds []string) error {
	r._batchIds = _batchIds
	r.Set("batch_ids", _batchIds)
	return nil
}

// GetBatchIds BatchIds Getter
func (r AlibabawdkfulfillbatchquerybybatchidsAPIRequest) GetBatchIds() []string {
	return r._batchIds
}

// SetNodeType is NodeType Setter
// 作业节点类型： WAREHOUSE：仓  DELIVERY_DOCK：配送站 SHOP：经营店
func (r *AlibabawdkfulfillbatchquerybybatchidsAPIRequest) SetNodeType(_nodeType string) error {
	r._nodeType = _nodeType
	r.Set("node_type", _nodeType)
	return nil
}

// GetNodeType NodeType Getter
func (r AlibabawdkfulfillbatchquerybybatchidsAPIRequest) GetNodeType() string {
	return r._nodeType
}

// SetNodeCode is NodeCode Setter
// warehouseCode, 出库仓，由基础店仓维护，盒马全域统一,
func (r *AlibabawdkfulfillbatchquerybybatchidsAPIRequest) SetNodeCode(_nodeCode string) error {
	r._nodeCode = _nodeCode
	r.Set("node_code", _nodeCode)
	return nil
}

// GetNodeCode NodeCode Getter
func (r AlibabawdkfulfillbatchquerybybatchidsAPIRequest) GetNodeCode() string {
	return r._nodeCode
}
