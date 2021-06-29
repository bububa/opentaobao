package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
精灵用户绑定第三方账号信息 API请求
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

// 初始化AlibabaAiUserQuickBindRequest对象
func NewAlibabaAiUserQuickBindRequest() *AlibabaAiUserQuickBindRequest{
    return &AlibabaAiUserQuickBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAiUserQuickBindRequest) GetApiMethodName() string {
    return "alibaba.ai.user.quick.bind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAiUserQuickBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SerialNo Setter
// 交易流水号（唯一即可，不参与业务运算）
func (r *AlibabaAiUserQuickBindRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaAiUserQuickBindRequest) GetSerialNo() string {
    return r.serialNo
}
// ExtUserType Setter
// 第三方用户类型
func (r *AlibabaAiUserQuickBindRequest) SetExtUserType(extUserType string) error {
    r.extUserType = extUserType
    r.Set("ext_user_type", extUserType)
    return nil
}

// ExtUserType Getter
func (r AlibabaAiUserQuickBindRequest) GetExtUserType() string {
    return r.extUserType
}
// ExtUserId Setter
// 第三用户ID
func (r *AlibabaAiUserQuickBindRequest) SetExtUserId(extUserId string) error {
    r.extUserId = extUserId
    r.Set("ext_user_id", extUserId)
    return nil
}

// ExtUserId Getter
func (r AlibabaAiUserQuickBindRequest) GetExtUserId() string {
    return r.extUserId
}
// ReqTime Setter
// 请求时间
func (r *AlibabaAiUserQuickBindRequest) SetReqTime(reqTime string) error {
    r.reqTime = reqTime
    r.Set("req_time", reqTime)
    return nil
}

// ReqTime Getter
func (r AlibabaAiUserQuickBindRequest) GetReqTime() string {
    return r.reqTime
}
// MerchantUserId Setter
// 商户的用户的唯一ID
func (r *AlibabaAiUserQuickBindRequest) SetMerchantUserId(merchantUserId string) error {
    r.merchantUserId = merchantUserId
    r.Set("merchant_user_id", merchantUserId)
    return nil
}

// MerchantUserId Getter
func (r AlibabaAiUserQuickBindRequest) GetMerchantUserId() string {
    return r.merchantUserId
}
// SchemaKey Setter
// 开放平台申请的schema
func (r *AlibabaAiUserQuickBindRequest) SetSchemaKey(schemaKey string) error {
    r.schemaKey = schemaKey
    r.Set("schema_key", schemaKey)
    return nil
}

// SchemaKey Getter
func (r AlibabaAiUserQuickBindRequest) GetSchemaKey() string {
    return r.schemaKey
}
