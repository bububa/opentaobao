package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单状态查询接口 APIRequest
alibaba.ele.enterprise.ordernew.getstatus

订单状态查询接口
*/
type AlibabaEleEnterpriseOrdernewGetstatusRequest struct {
    model.Params

    // 订单号
    elemeOrderId   string 

}

func NewAlibabaEleEnterpriseOrdernewGetstatusRequest() *AlibabaEleEnterpriseOrdernewGetstatusRequest{
    return &AlibabaEleEnterpriseOrdernewGetstatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseOrdernewGetstatusRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.getstatus"
}

func (r AlibabaEleEnterpriseOrdernewGetstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseOrdernewGetstatusRequest) SetElemeOrderId(elemeOrderId string) error {
    r.elemeOrderId = elemeOrderId
    r.Set("eleme_order_id", elemeOrderId)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewGetstatusRequest) GetElemeOrderId() string {
    return r.elemeOrderId
}

