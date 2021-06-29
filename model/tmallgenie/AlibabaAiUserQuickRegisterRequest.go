package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
精灵用户注册申请 API请求
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

// 初始化AlibabaAiUserQuickRegisterRequest对象
func NewAlibabaAiUserQuickRegisterRequest() *AlibabaAiUserQuickRegisterRequest{
    return &AlibabaAiUserQuickRegisterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAiUserQuickRegisterRequest) GetApiMethodName() string {
    return "alibaba.ai.user.quick.register"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAiUserQuickRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SerialNo Setter
// 请求交易流水号（唯一即可，不参与业务运算）
func (r *AlibabaAiUserQuickRegisterRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaAiUserQuickRegisterRequest) GetSerialNo() string {
    return r.serialNo
}
// ReqTime Setter
// 请求时间
func (r *AlibabaAiUserQuickRegisterRequest) SetReqTime(reqTime string) error {
    r.reqTime = reqTime
    r.Set("req_time", reqTime)
    return nil
}

// ReqTime Getter
func (r AlibabaAiUserQuickRegisterRequest) GetReqTime() string {
    return r.reqTime
}
// MerchantUserId Setter
// 商户的用户的唯一ID
func (r *AlibabaAiUserQuickRegisterRequest) SetMerchantUserId(merchantUserId string) error {
    r.merchantUserId = merchantUserId
    r.Set("merchant_user_id", merchantUserId)
    return nil
}

// MerchantUserId Getter
func (r AlibabaAiUserQuickRegisterRequest) GetMerchantUserId() string {
    return r.merchantUserId
}
// SchemaKey Setter
// 账户体系隔离
func (r *AlibabaAiUserQuickRegisterRequest) SetSchemaKey(schemaKey string) error {
    r.schemaKey = schemaKey
    r.Set("schema_key", schemaKey)
    return nil
}

// SchemaKey Getter
func (r AlibabaAiUserQuickRegisterRequest) GetSchemaKey() string {
    return r.schemaKey
}
