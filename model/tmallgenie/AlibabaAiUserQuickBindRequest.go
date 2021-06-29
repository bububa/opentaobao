package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
精灵用户绑定第三方账号信息 APIRequest
alibaba.ai.user.quick.bind

人工智能实验室精灵用户绑定第三方账号信息接口，开放给Iot厂商做为厂商上送第三方账号信息的接口
*/
type AlibabaAiUserQuickBindRequest struct {
    model.Params

    // 交易流水号（唯一即可，不参与业务运算）
    serialNo   string 

    // 第三方用户类型
    extUserType   string 

    // 第三用户ID
    extUserId   string 

    // 请求时间
    reqTime   string 

    // 商户的用户的唯一ID
    merchantUserId   string 

    // 开放平台申请的schema
    schemaKey   string 

}

func NewAlibabaAiUserQuickBindRequest() *AlibabaAiUserQuickBindRequest{
    return &AlibabaAiUserQuickBindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAiUserQuickBindRequest) GetApiMethodName() string {
    return "alibaba.ai.user.quick.bind"
}

func (r AlibabaAiUserQuickBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAiUserQuickBindRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

func (r AlibabaAiUserQuickBindRequest) GetSerialNo() string {
    return r.serialNo
}

func (r *AlibabaAiUserQuickBindRequest) SetExtUserType(extUserType string) error {
    r.extUserType = extUserType
    r.Set("ext_user_type", extUserType)
    return nil
}

func (r AlibabaAiUserQuickBindRequest) GetExtUserType() string {
    return r.extUserType
}

func (r *AlibabaAiUserQuickBindRequest) SetExtUserId(extUserId string) error {
    r.extUserId = extUserId
    r.Set("ext_user_id", extUserId)
    return nil
}

func (r AlibabaAiUserQuickBindRequest) GetExtUserId() string {
    return r.extUserId
}

func (r *AlibabaAiUserQuickBindRequest) SetReqTime(reqTime string) error {
    r.reqTime = reqTime
    r.Set("req_time", reqTime)
    return nil
}

func (r AlibabaAiUserQuickBindRequest) GetReqTime() string {
    return r.reqTime
}

func (r *AlibabaAiUserQuickBindRequest) SetMerchantUserId(merchantUserId string) error {
    r.merchantUserId = merchantUserId
    r.Set("merchant_user_id", merchantUserId)
    return nil
}

func (r AlibabaAiUserQuickBindRequest) GetMerchantUserId() string {
    return r.merchantUserId
}

func (r *AlibabaAiUserQuickBindRequest) SetSchemaKey(schemaKey string) error {
    r.schemaKey = schemaKey
    r.Set("schema_key", schemaKey)
    return nil
}

func (r AlibabaAiUserQuickBindRequest) GetSchemaKey() string {
    return r.schemaKey
}

