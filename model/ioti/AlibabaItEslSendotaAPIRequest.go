package ioti

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaItEslSendotaAPIRequest) Reset() {
	r._macAp = ""
	r._otaDataBase64String = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItEslSendotaAPIRequest) GetApiMethodName() string {
	return "alibaba.it.esl.sendota"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItEslSendotaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItEslSendotaAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMacAp is MacAp Setter
// mac
func (r *AlibabaItEslSendotaAPIRequest) SetMacAp(_macAp string) error {
	r._macAp = _macAp
	r.Set("mac_ap", _macAp)
	return nil
}

// GetMacAp MacAp Getter
func (r AlibabaItEslSendotaAPIRequest) GetMacAp() string {
	return r._macAp
}

// SetOtaDataBase64String is OtaDataBase64String Setter
// base64的ota包
func (r *AlibabaItEslSendotaAPIRequest) SetOtaDataBase64String(_otaDataBase64String string) error {
	r._otaDataBase64String = _otaDataBase64String
	r.Set("ota_data_base64_string", _otaDataBase64String)
	return nil
}

// GetOtaDataBase64String OtaDataBase64String Getter
func (r AlibabaItEslSendotaAPIRequest) GetOtaDataBase64String() string {
	return r._otaDataBase64String
}

var poolAlibabaItEslSendotaAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaItEslSendotaRequest()
	},
}

// GetAlibabaItEslSendotaRequest 从 sync.Pool 获取 AlibabaItEslSendotaAPIRequest
func GetAlibabaItEslSendotaAPIRequest() *AlibabaItEslSendotaAPIRequest {
	return poolAlibabaItEslSendotaAPIRequest.Get().(*AlibabaItEslSendotaAPIRequest)
}

// ReleaseAlibabaItEslSendotaAPIRequest 将 AlibabaItEslSendotaAPIRequest 放入 sync.Pool
func ReleaseAlibabaItEslSendotaAPIRequest(v *AlibabaItEslSendotaAPIRequest) {
	v.Reset()
	poolAlibabaItEslSendotaAPIRequest.Put(v)
}
