package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
签收接口 API请求
tmall.msf.receive

签收接口
*/
type TmallMsfReceiveRequest struct {
    model.Params
    // 1
    shopId   string
    // 1
    bizType   string
    // 1
    code   string
}

// 初始化TmallMsfReceiveRequest对象
func NewTmallMsfReceiveRequest() *TmallMsfReceiveRequest{
    return &TmallMsfReceiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMsfReceiveRequest) GetApiMethodName() string {
    return "tmall.msf.receive"
}

// IRequest interface 方法, 获取API参数
func (r TmallMsfReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 1
func (r *TmallMsfReceiveRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r TmallMsfReceiveRequest) GetShopId() string {
    return r.shopId
}
// BizType Setter
// 1
func (r *TmallMsfReceiveRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TmallMsfReceiveRequest) GetBizType() string {
    return r.bizType
}
// Code Setter
// 1
func (r *TmallMsfReceiveRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r TmallMsfReceiveRequest) GetCode() string {
    return r.code
}
