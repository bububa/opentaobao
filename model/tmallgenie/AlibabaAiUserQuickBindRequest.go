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
type AlibabaAiUserQuickBindAPIRequest struct {
    model.Params
    // 交易流水号（唯一即可，不参与业务运算）
    _serialNo   string
    // 第三方用户类型
    _extUserType   string
    // 第三用户ID
    _extUserId   string
    // 请求时间
    _reqTime   string
    // 商户的用户的唯一ID
    _merchantUserId   string
    // 开放平台申请的schema
    _schemaKey   string
}

// 初始化AlibabaAiUserQuickBindAPIRequest对象
func NewAlibabaAiUserQuickBindRequest() *AlibabaAiUserQuickBindAPIRequest{
    return &AlibabaAiUserQuickBindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAiUserQuickBindAPIRequest) GetApiMethodName() string {
    return "alibaba.ai.user.quick.bind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAiUserQuickBindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SerialNo Setter
// 交易流水号（唯一即可，不参与业务运算）
func (r *AlibabaAiUserQuickBindAPIRequest) SetSerialNo(_serialNo string) error {
    r._serialNo = _serialNo
    r.Set("serial_no", _serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaAiUserQuickBindAPIRequest) GetSerialNo() string {
    return r._serialNo
}
// ExtUserType Setter
// 第三方用户类型
func (r *AlibabaAiUserQuickBindAPIRequest) SetExtUserType(_extUserType string) error {
    r._extUserType = _extUserType
    r.Set("ext_user_type", _extUserType)
    return nil
}

// ExtUserType Getter
func (r AlibabaAiUserQuickBindAPIRequest) GetExtUserType() string {
    return r._extUserType
}
// ExtUserId Setter
// 第三用户ID
func (r *AlibabaAiUserQuickBindAPIRequest) SetExtUserId(_extUserId string) error {
    r._extUserId = _extUserId
    r.Set("ext_user_id", _extUserId)
    return nil
}

// ExtUserId Getter
func (r AlibabaAiUserQuickBindAPIRequest) GetExtUserId() string {
    return r._extUserId
}
// ReqTime Setter
// 请求时间
func (r *AlibabaAiUserQuickBindAPIRequest) SetReqTime(_reqTime string) error {
    r._reqTime = _reqTime
    r.Set("req_time", _reqTime)
    return nil
}

// ReqTime Getter
func (r AlibabaAiUserQuickBindAPIRequest) GetReqTime() string {
    return r._reqTime
}
// MerchantUserId Setter
// 商户的用户的唯一ID
func (r *AlibabaAiUserQuickBindAPIRequest) SetMerchantUserId(_merchantUserId string) error {
    r._merchantUserId = _merchantUserId
    r.Set("merchant_user_id", _merchantUserId)
    return nil
}

// MerchantUserId Getter
func (r AlibabaAiUserQuickBindAPIRequest) GetMerchantUserId() string {
    return r._merchantUserId
}
// SchemaKey Setter
// 开放平台申请的schema
func (r *AlibabaAiUserQuickBindAPIRequest) SetSchemaKey(_schemaKey string) error {
    r._schemaKey = _schemaKey
    r.Set("schema_key", _schemaKey)
    return nil
}

// SchemaKey Getter
func (r AlibabaAiUserQuickBindAPIRequest) GetSchemaKey() string {
    return r._schemaKey
}
