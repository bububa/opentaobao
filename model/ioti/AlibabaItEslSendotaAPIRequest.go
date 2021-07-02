package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItEslSendotaAPIRequest 电子价签ota接口 API请求
// alibaba.it.esl.sendota
//
// 厂测接口，电子价签ota接口
type AlibabaItEslSendotaAPIRequest struct {
	model.Params
	// mac
	_macAp string
	// base64的ota包
	_otaDataBase64String string
}

// NewAlibabaItEslSendotaRequest 初始化AlibabaItEslSendotaAPIRequest对象
func NewAlibabaItEslSendotaRequest() *AlibabaItEslSendotaAPIRequest {
	return &AlibabaItEslSendotaAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItEslSendotaAPIRequest) GetApiMethodName() string {
	return "alibaba.it.esl.sendota"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItEslSendotaAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MacAp Setter
// mac
func (r *AlibabaItEslSendotaAPIRequest) SetMacAp(_macAp string) error {
	r._macAp = _macAp
	r.Set("mac_ap", _macAp)
	return nil
}

// Get MacAp Getter
func (r AlibabaItEslSendotaAPIRequest) GetMacAp() string {
	return r._macAp
}

// Set is OtaDataBase64String Setter
// base64的ota包
func (r *AlibabaItEslSendotaAPIRequest) SetOtaDataBase64String(_otaDataBase64String string) error {
	r._otaDataBase64String = _otaDataBase64String
	r.Set("ota_data_base64_string", _otaDataBase64String)
	return nil
}

// Get OtaDataBase64String Getter
func (r AlibabaItEslSendotaAPIRequest) GetOtaDataBase64String() string {
	return r._otaDataBase64String
}
