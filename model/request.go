package model

import (
	"net/url"
	"time"
)

// APIFormat API请求及返回格式
type APIFormat = string

const (

	// JSON json api format
	JSON APIFormat = "json"

	// XML xml api format
	XML APIFormat = "xml"
)

// SignMethod API签名方法
type SignMethod = string

const (

	// MD5 md5 sign method
	MD5 SignMethod = "md5"

	// HMAC hmac sign method
	HMAC SignMethod = "hmac"
)

const (

	// DEFAULT_SIGN_METHOD 默认签名方法
	DEFAULT_SIGN_METHOD = MD5

	// DEFAULT_API_VERSION 默认API版本
	DEFAULT_API_VERSION = "2.0"

	// DEFAULT_API_FORMAT 默认api返回格式
	DEFAULT_API_FORMAT = JSON
)

// CommonRequest API请求通用参数
type CommonRequest struct {
	Method       string    `json:"method,omitempty"`         // API接口名称
	AppKey       string    `json:"app_key,omitempty"`        // TOP分配给应用的AppKey
	TargetAppKey string    `json:"target_app_key,omitempty"` // 被调用的目标AppKey，仅当被调用的API为第三方ISV提供时有效
	SignMethod   string    `json:"sign_method,omitempty"`    // 签名的摘要算法，可选值为：hmac，md5，hmac-sha256。
	Session      string    `json:"session,omitempty"`        // 用户登录授权成功后，TOP颁发给应用的授权信息，详细介绍请点击这里。当此API的标签上注明：“需要授权”，则此参数必传；“不需要授权”，则此参数不需要传；“可选授权”，则此参数为可选
	Timestamp    string    `json:"timestamp,omitempty"`      // 时间戳，格式为yyyy-MM-dd HH:mm:ss，时区为GMT+8，例如：2015-01-01 12:00:00。淘宝API服务端允许客户端请求最大时间误差为10分钟
	Format       APIFormat `json:"format,omitempty"`         // 响应格式。默认为xml格式，可选值：xml，json。
	V            string    `json:"v,omitempty"`              // API协议版本，可选值：2.0
	PartnerId    string    `json:"partner_id,omitempty"`     // 合作伙伴身份标识
	Simplify     bool      `json:"simplify,omitempty"`       // 是否采用精简JSON返回格式，仅当format=json时有效，默认值为：false
}

// NewCommonRequest 新建API请求通用参数
func NewCommonRequest(method string, appKey string) *CommonRequest {
	return &CommonRequest{
		Method:     method,
		AppKey:     appKey,
		SignMethod: DEFAULT_SIGN_METHOD,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Format:     DEFAULT_API_FORMAT,
		V:          DEFAULT_API_VERSION,
	}
}

// SetSession 设置access_session
func (c *CommonRequest) SetSession(session string) {
	c.Session = session
}

// SetAPIFormat 设置API格式
func (c *CommonRequest) SetAPIFormat(format APIFormat) {
	c.Format = format
}

// SetSignMethod 设置签名方法
func (c *CommonRequest) SetSignMethod(method SignMethod) {
	c.SignMethod = method
}

// GetParams 转换通用请求参数为map[string]string
func (c CommonRequest) GetParams() map[string]string {
	params := map[string]string{
		"method":      c.Method,
		"app_key":     c.AppKey,
		"sign_method": c.SignMethod,
		"timestamp":   c.Timestamp,
		"format":      c.Format,
		"v":           c.V,
	}
	if c.TargetAppKey != "" {
		params["target_app_key"] = c.TargetAppKey
	}
	if c.Session != "" {
		params["session"] = c.Session
	}
	if c.PartnerId != "" {
		params["partner_id"] = c.PartnerId
	}
	if c.Format == JSON {
		params["simplify"] = "true"
	}

	return params
}

// IRequest is API Request interace
type IRequest interface {
	GetApiMethodName() string
	GetApiParams() url.Values
	NeedMultipart() bool
	GetRawParams() map[string]*ParamValue
}
