package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠规则删除 API请求
alibaba.asr.dataservice.promotionrule.delete

删除优惠规则，例如星巴克删除优惠规则
*/
type AlibabaAsrDataservicePromotionruleDeleteRequest struct {
    model.Params
    // poskey
    _posKey   int64
}

// 初始化AlibabaAsrDataservicePromotionruleDeleteRequest对象
func NewAlibabaAsrDataservicePromotionruleDeleteRequest() *AlibabaAsrDataservicePromotionruleDeleteRequest{
    return &AlibabaAsrDataservicePromotionruleDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAsrDataservicePromotionruleDeleteRequest) GetApiMethodName() string {
    return "alibaba.asr.dataservice.promotionrule.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAsrDataservicePromotionruleDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosKey Setter
// poskey
func (r *AlibabaAsrDataservicePromotionruleDeleteRequest) SetPosKey(_posKey int64) error {
    r._posKey = _posKey
    r.Set("pos_key", _posKey)
    return nil
}

// PosKey Getter
func (r AlibabaAsrDataservicePromotionruleDeleteRequest) GetPosKey() int64 {
    return r._posKey
}
