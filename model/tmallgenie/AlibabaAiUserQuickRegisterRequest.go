package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
精灵用户注册申请 APIRequest
alibaba.ai.user.quick.register

人工智能实验室精灵用户注册申请接口，开放给Iot厂商做厂商会员数据上报
*/
type AlibabaAiUserQuickRegisterRequest struct {
    model.Params

    // 请求交易流水号（唯一即可，不参与业务运算）
    serialNo   string 

    // 请求时间
    reqTime   string 

    // 商户的用户的唯一ID
    merchantUserId   string 

    // 账户体系隔离
    schemaKey   string 

}

func NewAlibabaAiUserQuickRegisterRequest() *AlibabaAiUserQuickRegisterRequest{
    return &AlibabaAiUserQuickRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAiUserQuickRegisterRequest) GetApiMethodName() string {
    return "alibaba.ai.user.quick.register"
}

func (r AlibabaAiUserQuickRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAiUserQuickRegisterRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

func (r AlibabaAiUserQuickRegisterRequest) GetSerialNo() string {
    return r.serialNo
}

func (r *AlibabaAiUserQuickRegisterRequest) SetReqTime(reqTime string) error {
    r.reqTime = reqTime
    r.Set("req_time", reqTime)
    return nil
}

func (r AlibabaAiUserQuickRegisterRequest) GetReqTime() string {
    return r.reqTime
}

func (r *AlibabaAiUserQuickRegisterRequest) SetMerchantUserId(merchantUserId string) error {
    r.merchantUserId = merchantUserId
    r.Set("merchant_user_id", merchantUserId)
    return nil
}

func (r AlibabaAiUserQuickRegisterRequest) GetMerchantUserId() string {
    return r.merchantUserId
}

func (r *AlibabaAiUserQuickRegisterRequest) SetSchemaKey(schemaKey string) error {
    r.schemaKey = schemaKey
    r.Set("schema_key", schemaKey)
    return nil
}

func (r AlibabaAiUserQuickRegisterRequest) GetSchemaKey() string {
    return r.schemaKey
}

