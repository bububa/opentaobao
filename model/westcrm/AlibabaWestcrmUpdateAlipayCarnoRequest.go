package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新支付宝业务卡号 APIRequest
alibaba.westcrm.update.alipay.carno

更新支付宝业务卡号
*/
type AlibabaWestcrmUpdateAlipayCarnoRequest struct {
    model.Params

    // 商场id
    mallId   int64 

    // 用户id
    id   int64 

    // 2088102011918821
    alipayCardNo   string 

    // appkey
    westcrmAppKey   string 

}

func NewAlibabaWestcrmUpdateAlipayCarnoRequest() *AlibabaWestcrmUpdateAlipayCarnoRequest{
    return &AlibabaWestcrmUpdateAlipayCarnoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetApiMethodName() string {
    return "alibaba.westcrm.update.alipay.carno"
}

func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmUpdateAlipayCarnoRequest) SetMallId(mallId int64) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetMallId() int64 {
    return r.mallId
}

func (r *AlibabaWestcrmUpdateAlipayCarnoRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetId() int64 {
    return r.id
}

func (r *AlibabaWestcrmUpdateAlipayCarnoRequest) SetAlipayCardNo(alipayCardNo string) error {
    r.alipayCardNo = alipayCardNo
    r.Set("alipay_card_no", alipayCardNo)
    return nil
}

func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetAlipayCardNo() string {
    return r.alipayCardNo
}

func (r *AlibabaWestcrmUpdateAlipayCarnoRequest) SetWestcrmAppKey(westcrmAppKey string) error {
    r.westcrmAppKey = westcrmAppKey
    r.Set("westcrm_app_key", westcrmAppKey)
    return nil
}

func (r AlibabaWestcrmUpdateAlipayCarnoRequest) GetWestcrmAppKey() string {
    return r.westcrmAppKey
}

