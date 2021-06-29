package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价签设备信息查询接口 APIRequest
taobao.uscesl.biz.esl.info.get

价签设备信息查询接口
*/
type TaobaoUsceslBizEslInfoGetRequest struct {
    model.Params

    // 价签条码
    eslBarCode   string 

    // 门店storeId
    storeId   int64 

    // 商家ID
    bizBrandKey   string 

}

func NewTaobaoUsceslBizEslInfoGetRequest() *TaobaoUsceslBizEslInfoGetRequest{
    return &TaobaoUsceslBizEslInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsceslBizEslInfoGetRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.esl.info.get"
}

func (r TaobaoUsceslBizEslInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsceslBizEslInfoGetRequest) SetEslBarCode(eslBarCode string) error {
    r.eslBarCode = eslBarCode
    r.Set("esl_bar_code", eslBarCode)
    return nil
}

func (r TaobaoUsceslBizEslInfoGetRequest) GetEslBarCode() string {
    return r.eslBarCode
}

func (r *TaobaoUsceslBizEslInfoGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoUsceslBizEslInfoGetRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoUsceslBizEslInfoGetRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

func (r TaobaoUsceslBizEslInfoGetRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}

