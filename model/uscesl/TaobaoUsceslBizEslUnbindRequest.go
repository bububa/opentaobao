package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子价签解绑接口 API请求
taobao.uscesl.biz.esl.unbind

电子价签解绑接口
*/
type TaobaoUsceslBizEslUnbindRequest struct {
    model.Params
    // 价签条码
    eslBarCode   string
    // 价签系统注册的门店storeId
    storeId   int64
    // 商家标识key
    bizBrandKey   string
}

// 初始化TaobaoUsceslBizEslUnbindRequest对象
func NewTaobaoUsceslBizEslUnbindRequest() *TaobaoUsceslBizEslUnbindRequest{
    return &TaobaoUsceslBizEslUnbindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizEslUnbindRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.esl.unbind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizEslUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EslBarCode Setter
// 价签条码
func (r *TaobaoUsceslBizEslUnbindRequest) SetEslBarCode(eslBarCode string) error {
    r.eslBarCode = eslBarCode
    r.Set("esl_bar_code", eslBarCode)
    return nil
}

// EslBarCode Getter
func (r TaobaoUsceslBizEslUnbindRequest) GetEslBarCode() string {
    return r.eslBarCode
}
// StoreId Setter
// 价签系统注册的门店storeId
func (r *TaobaoUsceslBizEslUnbindRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizEslUnbindRequest) GetStoreId() int64 {
    return r.storeId
}
// BizBrandKey Setter
// 商家标识key
func (r *TaobaoUsceslBizEslUnbindRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizEslUnbindRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}
