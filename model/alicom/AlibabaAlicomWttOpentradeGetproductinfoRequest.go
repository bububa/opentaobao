package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询存送产品信息 API请求
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

// 初始化AlibabaAlicomWttOpentradeGetproductinfoRequest对象
func NewAlibabaAlicomWttOpentradeGetproductinfoRequest() *AlibabaAlicomWttOpentradeGetproductinfoRequest{
    return &AlibabaAlicomWttOpentradeGetproductinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlicomWttOpentradeGetproductinfoRequest) GetApiMethodName() string {
    return "alibaba.alicom.wtt.opentrade.getproductinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlicomWttOpentradeGetproductinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 阿里通信产品ID
func (r *AlibabaAlicomWttOpentradeGetproductinfoRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AlibabaAlicomWttOpentradeGetproductinfoRequest) GetProductId() string {
    return r.productId
}
// BizType Setter
// 类型
func (r *AlibabaAlicomWttOpentradeGetproductinfoRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r AlibabaAlicomWttOpentradeGetproductinfoRequest) GetBizType() string {
    return r.bizType
}
