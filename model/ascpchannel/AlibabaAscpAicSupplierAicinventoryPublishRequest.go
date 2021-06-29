package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家仓操作aic库存发布服务 APIRequest
alibaba.ascp.aic.supplier.aicinventory.publish

商家调用这个接口来发布增加库存数据
*/
type AlibabaAscpAicSupplierAicinventoryPublishRequest struct {
    model.Params

    // 库存发布请求参数
    aicInventoryPublishRequest   *Aicinventorypublishrequest 

}

func NewAlibabaAscpAicSupplierAicinventoryPublishRequest() *AlibabaAscpAicSupplierAicinventoryPublishRequest{
    return &AlibabaAscpAicSupplierAicinventoryPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpAicSupplierAicinventoryPublishRequest) GetApiMethodName() string {
    return "alibaba.ascp.aic.supplier.aicinventory.publish"
}

func (r AlibabaAscpAicSupplierAicinventoryPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpAicSupplierAicinventoryPublishRequest) SetAicInventoryPublishRequest(aicInventoryPublishRequest *Aicinventorypublishrequest) error {
    r.aicInventoryPublishRequest = aicInventoryPublishRequest
    r.Set("aic_inventory_publish_request", aicInventoryPublishRequest)
    return nil
}

func (r AlibabaAscpAicSupplierAicinventoryPublishRequest) GetAicInventoryPublishRequest() *Aicinventorypublishrequest {
    return r.aicInventoryPublishRequest
}

