package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅核销接口 API请求
tmall.msf.verify

msf服务核销的top接口
*/
type TmallMsfVerifyRequest struct {
    model.Params
    // 111
    _shopId   string
    // 111
    _bizType   string
    // 111
    _code   string
}

// 初始化TmallMsfVerifyRequest对象
func NewTmallMsfVerifyRequest() *TmallMsfVerifyRequest{
    return &TmallMsfVerifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMsfVerifyRequest) GetApiMethodName() string {
    return "tmall.msf.verify"
}

// IRequest interface 方法, 获取API参数
func (r TmallMsfVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 111
func (r *TmallMsfVerifyRequest) SetShopId(_shopId string) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r TmallMsfVerifyRequest) GetShopId() string {
    return r._shopId
}
// BizType Setter
// 111
func (r *TmallMsfVerifyRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TmallMsfVerifyRequest) GetBizType() string {
    return r._bizType
}
// Code Setter
// 111
func (r *TmallMsfVerifyRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TmallMsfVerifyRequest) GetCode() string {
    return r._code
}
