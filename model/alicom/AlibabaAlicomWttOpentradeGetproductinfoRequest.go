package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询存送产品信息 APIRequest
alibaba.alicom.wtt.opentrade.getproductinfo

话费宝查询产品信息相关配置
*/
type AlibabaAlicomWttOpentradeGetproductinfoRequest struct {
    model.Params

    // 阿里通信产品ID
    productId   string 

    // 类型
    bizType   string 

}

func NewAlibabaAlicomWttOpentradeGetproductinfoRequest() *AlibabaAlicomWttOpentradeGetproductinfoRequest{
    return &AlibabaAlicomWttOpentradeGetproductinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlicomWttOpentradeGetproductinfoRequest) GetApiMethodName() string {
    return "alibaba.alicom.wtt.opentrade.getproductinfo"
}

func (r AlibabaAlicomWttOpentradeGetproductinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlicomWttOpentradeGetproductinfoRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AlibabaAlicomWttOpentradeGetproductinfoRequest) GetProductId() string {
    return r.productId
}

func (r *AlibabaAlicomWttOpentradeGetproductinfoRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r AlibabaAlicomWttOpentradeGetproductinfoRequest) GetBizType() string {
    return r.bizType
}

