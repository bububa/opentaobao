package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知制卡商制卡 APIRequest
alibaba.fundplatform.cardorder.make

该接口由内部定义，外部制卡商实现。当需要制卡商进行制卡操作时，将会调用该接口。
*/
type AlibabaFundplatformCardorderMakeRequest struct {
    model.Params

    // 物流信息
    logisticsInfo   *AlibabaFundplatformCardorderMakeStruct 

    // 卡模板信息列表
    cardProductInfos   []AlibabaFundplatformCardorderMakeStruct 

    // 子制卡单ID
    cardOrderId   int64 

    // 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
    ownSign   string 

}

func NewAlibabaFundplatformCardorderMakeRequest() *AlibabaFundplatformCardorderMakeRequest{
    return &AlibabaFundplatformCardorderMakeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformCardorderMakeRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorder.make"
}

func (r AlibabaFundplatformCardorderMakeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformCardorderMakeRequest) SetLogisticsInfo(logisticsInfo *AlibabaFundplatformCardorderMakeStruct) error {
    r.logisticsInfo = logisticsInfo
    r.Set("logistics_info", logisticsInfo)
    return nil
}

func (r AlibabaFundplatformCardorderMakeRequest) GetLogisticsInfo() *AlibabaFundplatformCardorderMakeStruct {
    return r.logisticsInfo
}

func (r *AlibabaFundplatformCardorderMakeRequest) SetCardProductInfos(cardProductInfos []AlibabaFundplatformCardorderMakeStruct) error {
    r.cardProductInfos = cardProductInfos
    r.Set("card_product_infos", cardProductInfos)
    return nil
}

func (r AlibabaFundplatformCardorderMakeRequest) GetCardProductInfos() []AlibabaFundplatformCardorderMakeStruct {
    return r.cardProductInfos
}

func (r *AlibabaFundplatformCardorderMakeRequest) SetCardOrderId(cardOrderId int64) error {
    r.cardOrderId = cardOrderId
    r.Set("card_order_id", cardOrderId)
    return nil
}

func (r AlibabaFundplatformCardorderMakeRequest) GetCardOrderId() int64 {
    return r.cardOrderId
}

func (r *AlibabaFundplatformCardorderMakeRequest) SetOwnSign(ownSign string) error {
    r.ownSign = ownSign
    r.Set("own_sign", ownSign)
    return nil
}

func (r AlibabaFundplatformCardorderMakeRequest) GetOwnSign() string {
    return r.ownSign
}

