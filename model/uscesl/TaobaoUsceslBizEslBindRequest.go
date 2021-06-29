package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子价签绑定接口 APIRequest
taobao.uscesl.biz.esl.bind

电子价签商品绑定接口
*/
type TaobaoUsceslBizEslBindRequest struct {
    model.Params

    // 价签条码
    eslBarCode   string 

    // 商品条码
    itemBarCode   string 

    // 门店storeId
    storeId   int64 

    // 商家ID
    bizBrandKey   string 

    // 额外扩展信息
    extendInfo   string 

}

func NewTaobaoUsceslBizEslBindRequest() *TaobaoUsceslBizEslBindRequest{
    return &TaobaoUsceslBizEslBindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsceslBizEslBindRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.esl.bind"
}

func (r TaobaoUsceslBizEslBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsceslBizEslBindRequest) SetEslBarCode(eslBarCode string) error {
    r.eslBarCode = eslBarCode
    r.Set("esl_bar_code", eslBarCode)
    return nil
}

func (r TaobaoUsceslBizEslBindRequest) GetEslBarCode() string {
    return r.eslBarCode
}

func (r *TaobaoUsceslBizEslBindRequest) SetItemBarCode(itemBarCode string) error {
    r.itemBarCode = itemBarCode
    r.Set("item_bar_code", itemBarCode)
    return nil
}

func (r TaobaoUsceslBizEslBindRequest) GetItemBarCode() string {
    return r.itemBarCode
}

func (r *TaobaoUsceslBizEslBindRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoUsceslBizEslBindRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoUsceslBizEslBindRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

func (r TaobaoUsceslBizEslBindRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}

func (r *TaobaoUsceslBizEslBindRequest) SetExtendInfo(extendInfo string) error {
    r.extendInfo = extendInfo
    r.Set("extend_info", extendInfo)
    return nil
}

func (r TaobaoUsceslBizEslBindRequest) GetExtendInfo() string {
    return r.extendInfo
}

