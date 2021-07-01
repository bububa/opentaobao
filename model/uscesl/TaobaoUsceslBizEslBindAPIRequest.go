package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子价签绑定接口 API请求
taobao.uscesl.biz.esl.bind

电子价签商品绑定接口
*/
type TaobaoUsceslBizEslBindAPIRequest struct {
    model.Params
    // 价签条码
    _eslBarCode   string
    // 商品条码
    _itemBarCode   string
    // 门店storeId
    _storeId   int64
    // 商家ID
    _bizBrandKey   string
    // 额外扩展信息
    _extendInfo   string
}

// 初始化TaobaoUsceslBizEslBindAPIRequest对象
func NewTaobaoUsceslBizEslBindRequest() *TaobaoUsceslBizEslBindAPIRequest{
    return &TaobaoUsceslBizEslBindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizEslBindAPIRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.esl.bind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizEslBindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EslBarCode Setter
// 价签条码
func (r *TaobaoUsceslBizEslBindAPIRequest) SetEslBarCode(_eslBarCode string) error {
    r._eslBarCode = _eslBarCode
    r.Set("esl_bar_code", _eslBarCode)
    return nil
}

// EslBarCode Getter
func (r TaobaoUsceslBizEslBindAPIRequest) GetEslBarCode() string {
    return r._eslBarCode
}
// ItemBarCode Setter
// 商品条码
func (r *TaobaoUsceslBizEslBindAPIRequest) SetItemBarCode(_itemBarCode string) error {
    r._itemBarCode = _itemBarCode
    r.Set("item_bar_code", _itemBarCode)
    return nil
}

// ItemBarCode Getter
func (r TaobaoUsceslBizEslBindAPIRequest) GetItemBarCode() string {
    return r._itemBarCode
}
// StoreId Setter
// 门店storeId
func (r *TaobaoUsceslBizEslBindAPIRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizEslBindAPIRequest) GetStoreId() int64 {
    return r._storeId
}
// BizBrandKey Setter
// 商家ID
func (r *TaobaoUsceslBizEslBindAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
    r._bizBrandKey = _bizBrandKey
    r.Set("biz_brand_key", _bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizEslBindAPIRequest) GetBizBrandKey() string {
    return r._bizBrandKey
}
// ExtendInfo Setter
// 额外扩展信息
func (r *TaobaoUsceslBizEslBindAPIRequest) SetExtendInfo(_extendInfo string) error {
    r._extendInfo = _extendInfo
    r.Set("extend_info", _extendInfo)
    return nil
}

// ExtendInfo Getter
func (r TaobaoUsceslBizEslBindAPIRequest) GetExtendInfo() string {
    return r._extendInfo
}
