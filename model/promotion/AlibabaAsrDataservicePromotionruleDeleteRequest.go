package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠规则删除 APIRequest
alibaba.asr.dataservice.promotionrule.delete

删除优惠规则，例如星巴克删除优惠规则
*/
type AlibabaAsrDataservicePromotionruleDeleteRequest struct {
    model.Params

    // poskey
    posKey   int64 

}

func NewAlibabaAsrDataservicePromotionruleDeleteRequest() *AlibabaAsrDataservicePromotionruleDeleteRequest{
    return &AlibabaAsrDataservicePromotionruleDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAsrDataservicePromotionruleDeleteRequest) GetApiMethodName() string {
    return "alibaba.asr.dataservice.promotionrule.delete"
}

func (r AlibabaAsrDataservicePromotionruleDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAsrDataservicePromotionruleDeleteRequest) SetPosKey(posKey int64) error {
    r.posKey = posKey
    r.Set("pos_key", posKey)
    return nil
}

func (r AlibabaAsrDataservicePromotionruleDeleteRequest) GetPosKey() int64 {
    return r.posKey
}

