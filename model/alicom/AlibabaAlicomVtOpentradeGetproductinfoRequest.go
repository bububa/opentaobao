package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询新虚拟产品配置信息 API请求
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

// 初始化AlibabaAlicomVtOpentradeGetproductinfoRequest对象
func NewAlibabaAlicomVtOpentradeGetproductinfoRequest() *AlibabaAlicomVtOpentradeGetproductinfoRequest{
    return &AlibabaAlicomVtOpentradeGetproductinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlicomVtOpentradeGetproductinfoRequest) GetApiMethodName() string {
    return "alibaba.alicom.vt.opentrade.getproductinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlicomVtOpentradeGetproductinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 阿里通信产品ID
func (r *AlibabaAlicomVtOpentradeGetproductinfoRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AlibabaAlicomVtOpentradeGetproductinfoRequest) GetProductId() string {
    return r.productId
}
// BizType Setter
// 类型
func (r *AlibabaAlicomVtOpentradeGetproductinfoRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r AlibabaAlicomVtOpentradeGetproductinfoRequest) GetBizType() string {
    return r.bizType
}
