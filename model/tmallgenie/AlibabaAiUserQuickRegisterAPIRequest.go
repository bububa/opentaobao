package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiUserQuickRegisterAPIRequest 精灵用户注册申请 API请求
// alibaba.ai.user.quick.register
//
// 人工智能实验室精灵用户注册申请接口，开放给Iot厂商做厂商会员数据上报
type AlibabaAiUserQuickRegisterAPIRequest struct {
	model.Params
	// 请求交易流水号（唯一即可，不参与业务运算）
	_serialNo string
	// 请求时间
	_reqTime string
	// 商户的用户的唯一ID
	_merchantUserId string
	// 账户体系隔离
	_schemaKey string
}

// NewAlibabaAiUserQuickRegisterRequest 初始化AlibabaAiUserQuickRegisterAPIRequest对象
func NewAlibabaAiUserQuickRegisterRequest() *AlibabaAiUserQuickRegisterAPIRequest {
	return &AlibabaAiUserQuickRegisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAiUserQuickRegisterAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.user.quick.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAiUserQuickRegisterAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSerialNo is SerialNo Setter
// 请求交易流水号（唯一即可，不参与业务运算）
func (r *AlibabaAiUserQuickRegisterAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaAiUserQuickRegisterAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// SetReqTime is ReqTime Setter
// 请求时间
func (r *AlibabaAiUserQuickRegisterAPIRequest) SetReqTime(_reqTime string) error {
	r._reqTime = _reqTime
	r.Set("req_time", _reqTime)
	return nil
}

// GetReqTime ReqTime Getter
func (r AlibabaAiUserQuickRegisterAPIRequest) GetReqTime() string {
	return r._reqTime
}

// SetMerchantUserId is MerchantUserId Setter
// 商户的用户的唯一ID
func (r *AlibabaAiUserQuickRegisterAPIRequest) SetMerchantUserId(_merchantUserId string) error {
	r._merchantUserId = _merchantUserId
	r.Set("merchant_user_id", _merchantUserId)
	return nil
}

// GetMerchantUserId MerchantUserId Getter
func (r AlibabaAiUserQuickRegisterAPIRequest) GetMerchantUserId() string {
	return r._merchantUserId
}

// SetSchemaKey is SchemaKey Setter
// 账户体系隔离
func (r *AlibabaAiUserQuickRegisterAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// GetSchemaKey SchemaKey Getter
func (r AlibabaAiUserQuickRegisterAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}
