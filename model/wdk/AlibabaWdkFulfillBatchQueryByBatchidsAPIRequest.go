package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
作业小票查询接口 API请求
alibaba.wdk.fulfill.batch.query.by.batchids

根据节点等条件查询履约单小票信息
*/
type AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest struct {
    model.Params
    // 作业节点类型： WAREHOUSE：仓  DELIVERY_DOCK：配送站 SHOP：经营店
    _nodeType   string
    // warehouseCode, 出库仓，由基础店仓维护，盒马全域统一,
    _nodeCode   string
    // 批次号列表
    _batchIds   []string
}

// 初始化AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest对象
func NewAlibabaWdkFulfillBatchQueryByBatchidsRequest() *AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest{
    return &AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.batch.query.by.batchids"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NodeType Setter
// 作业节点类型： WAREHOUSE：仓  DELIVERY_DOCK：配送站 SHOP：经营店
func (r *AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) SetNodeType(_nodeType string) error {
    r._nodeType = _nodeType
    r.Set("node_type", _nodeType)
    return nil
}

// NodeType Getter
func (r AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) GetNodeType() string {
    return r._nodeType
}
// NodeCode Setter
// warehouseCode, 出库仓，由基础店仓维护，盒马全域统一,
func (r *AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) SetNodeCode(_nodeCode string) error {
    r._nodeCode = _nodeCode
    r.Set("node_code", _nodeCode)
    return nil
}

// NodeCode Getter
func (r AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) GetNodeCode() string {
    return r._nodeCode
}
// BatchIds Setter
// 批次号列表
func (r *AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) SetBatchIds(_batchIds []string) error {
    r._batchIds = _batchIds
    r.Set("batch_ids", _batchIds)
    return nil
}

// BatchIds Getter
func (r AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest) GetBatchIds() []string {
    return r._batchIds
}
