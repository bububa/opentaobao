package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaiteslsendotaAPIRequest 电子价签ota接口 API请求
// alibaba.it.esl.sendota
//
// 厂测接口，电子价签ota接口
type AlibabaiteslsendotaAPIRequest struct {
	model.Params
	// mac
	_macAp string
	// base64的ota包
	_otaDataBase64String string
}

// NewAlibabaiteslsendotaRequest 初始化AlibabaiteslsendotaAPIRequest对象
func NewAlibabaiteslsendotaRequest() *AlibabaiteslsendotaAPIRequest {
	return &AlibabaiteslsendotaAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaiteslsendotaAPIRequest) GetApiMethodName() string {
	return "alibaba.it.esl.sendota"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaiteslsendotaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaiteslsendotaAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMacAp is MacAp Setter
// mac
func (r *AlibabaiteslsendotaAPIRequest) SetMacAp(_macAp string) error {
	r._macAp = _macAp
	r.Set("mac_ap", _macAp)
	return nil
}

// GetMacAp MacAp Getter
func (r AlibabaiteslsendotaAPIRequest) GetMacAp() string {
	return r._macAp
}

// SetOtaDataBase64String is OtaDataBase64String Setter
// base64的ota包
func (r *AlibabaiteslsendotaAPIRequest) SetOtaDataBase64String(_otaDataBase64String string) error {
	r._otaDataBase64String = _otaDataBase64String
	r.Set("ota_data_base64_string", _otaDataBase64String)
	return nil
}

// GetOtaDataBase64String OtaDataBase64String Getter
func (r AlibabaiteslsendotaAPIRequest) GetOtaDataBase64String() string {
	return r._otaDataBase64String
}
