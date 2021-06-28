package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
作业小票查询接口 APIRequest
alibaba.wdk.fulfill.batch.query.by.batchids

根据节点等条件查询履约单小票信息
*/
type AlibabaWdkFulfillBatchQueryByBatchidsRequest struct {
    model.Params

    // 作业节点类型： WAREHOUSE：仓  DELIVERY_DOCK：配送站 SHOP：经营店
    nodeType   string 

    // warehouseCode, 出库仓，由基础店仓维护，盒马全域统一,
    nodeCode   string 

    // 批次号列表
    batchIds   []string 

}

func NewAlibabaWdkFulfillBatchQueryByBatchidsRequest() *AlibabaWdkFulfillBatchQueryByBatchidsRequest{
    return &AlibabaWdkFulfillBatchQueryByBatchidsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkFulfillBatchQueryByBatchidsRequest) GetApiMethodName() string {
    return "alibaba.wdk.fulfill.batch.query.by.batchids"
}

func (r AlibabaWdkFulfillBatchQueryByBatchidsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkFulfillBatchQueryByBatchidsRequest) SetNodeType(nodeType string) error {
    r.nodeType = nodeType
    r.Set("node_type", nodeType)
    return nil
}

func (r AlibabaWdkFulfillBatchQueryByBatchidsRequest) GetNodeType() string {
    return r.nodeType
}

func (r *AlibabaWdkFulfillBatchQueryByBatchidsRequest) SetNodeCode(nodeCode string) error {
    r.nodeCode = nodeCode
    r.Set("node_code", nodeCode)
    return nil
}

func (r AlibabaWdkFulfillBatchQueryByBatchidsRequest) GetNodeCode() string {
    return r.nodeCode
}

func (r *AlibabaWdkFulfillBatchQueryByBatchidsRequest) SetBatchIds(batchIds []string) error {
    r.batchIds = batchIds
    r.Set("batch_ids", batchIds)
    return nil
}

func (r AlibabaWdkFulfillBatchQueryByBatchidsRequest) GetBatchIds() []string {
    return r.batchIds
}

