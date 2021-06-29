package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商直发-商家仓库存查询服务 APIRequest
alibaba.ascp.aic.supplier.aicinventory.channel.inventory.query

提供商家基于货品、供应商、仓，查询ascp 实时商家仓库存查询数据。
*/
type AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryRequest struct {
    model.Params

    // 商家仓库存查询请求参数
    merchantInventoryQueryRequest   *MerchantInventoryQuery 

}

func NewAlibabaAscpAicSupplierAicinventoryChannelInventoryQueryRequest() *AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryRequest{
    return &AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.aic.supplier.aicinventory.channel.inventory.query"
}

func (r AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryRequest) SetMerchantInventoryQueryRequest(merchantInventoryQueryRequest *MerchantInventoryQuery) error {
    r.merchantInventoryQueryRequest = merchantInventoryQueryRequest
    r.Set("merchant_inventory_query_request", merchantInventoryQueryRequest)
    return nil
}

func (r AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryRequest) GetMerchantInventoryQueryRequest() *MerchantInventoryQuery {
    return r.merchantInventoryQueryRequest
}

