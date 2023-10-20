package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqappshieldAPIRequest 应用加固接口 API请求
// alibaba.security.jaq.app.shield
//
// 提交应用进行应用加固,加固后需通过alibaba.security.jaq.app.shieldresult.get接口查询加固结果
type AlibabasecurityjaqappshieldAPIRequest struct {
	model.Params
	// 待加固的应用信息
	_appInfo *ScanAppInfo
	// 渠道列表,多渠道加固时填写
	_channel *ShieldChannel
}

// NewAlibabasecurityjaqappshieldRequest 初始化AlibabasecurityjaqappshieldAPIRequest对象
func NewAlibabasecurityjaqappshieldRequest() *AlibabasecurityjaqappshieldAPIRequest {
	return &AlibabasecurityjaqappshieldAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqappshieldAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.shield"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqappshieldAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqappshieldAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppInfo is AppInfo Setter
// 待加固的应用信息
func (r *AlibabasecurityjaqappshieldAPIRequest) SetAppInfo(_appInfo *ScanAppInfo) error {
	r._appInfo = _appInfo
	r.Set("app_info", _appInfo)
	return nil
}

// GetAppInfo AppInfo Getter
func (r AlibabasecurityjaqappshieldAPIRequest) GetAppInfo() *ScanAppInfo {
	return r._appInfo
}

// SetChannel is Channel Setter
// 渠道列表,多渠道加固时填写
func (r *AlibabasecurityjaqappshieldAPIRequest) SetChannel(_channel *ShieldChannel) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r AlibabasecurityjaqappshieldAPIRequest) GetChannel() *ShieldChannel {
	return r._channel
}
