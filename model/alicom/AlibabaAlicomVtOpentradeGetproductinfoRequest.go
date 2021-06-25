package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询新虚拟产品配置信息 APIRequest
alibaba.alicom.vt.opentrade.getproductinfo

话费宝查询产品信息相关配置
*/
type AlibabaAlicomVtOpentradeGetproductinfoRequest struct {
    model.Params

    // 阿里通信产品ID
    productId   string 

    // 类型
    bizType   string 

}

func NewAlibabaAlicomVtOpentradeGetproductinfoRequest() *AlibabaAlicomVtOpentradeGetproductinfoRequest{
    return &AlibabaAlicomVtOpentradeGetproductinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlicomVtOpentradeGetproductinfoRequest) GetApiMethodName() string {
    return "alibaba.alicom.vt.opentrade.getproductinfo"
}

func (r AlibabaAlicomVtOpentradeGetproductinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlicomVtOpentradeGetproductinfoRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AlibabaAlicomVtOpentradeGetproductinfoRequest) GetProductId() string {
    return r.productId
}

func (r *AlibabaAlicomVtOpentradeGetproductinfoRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r AlibabaAlicomVtOpentradeGetproductinfoRequest) GetBizType() string {
    return r.bizType
}

