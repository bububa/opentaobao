package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
签收接口 APIRequest
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

func NewTmallMsfReceiveRequest() *TmallMsfReceiveRequest{
    return &TmallMsfReceiveRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMsfReceiveRequest) GetApiMethodName() string {
    return "tmall.msf.receive"
}

func (r TmallMsfReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMsfReceiveRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r TmallMsfReceiveRequest) GetShopId() string {
    return r.shopId
}

func (r *TmallMsfReceiveRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TmallMsfReceiveRequest) GetBizType() string {
    return r.bizType
}

func (r *TmallMsfReceiveRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TmallMsfReceiveRequest) GetCode() string {
    return r.code
}

