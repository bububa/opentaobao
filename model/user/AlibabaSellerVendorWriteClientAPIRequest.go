package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSellerVendorWriteClientAPIRequest 客户动态回写 API请求
// alibaba.seller.vendor.write.client
//
// 客户动态开放API接口，外部服务商回写数据
type AlibabaSellerVendorWriteClientAPIRequest struct {
	model.Params
	// 开放平台appId
	_appId string
	// 要回写的数据
	_paramThirdPartyClientDataParams *ThirdPartyClientDataParams
}

// NewAlibabaSellerVendorWriteClientRequest 初始化AlibabaSellerVendorWriteClientAPIRequest对象
func NewAlibabaSellerVendorWriteClientRequest() *AlibabaSellerVendorWriteClientAPIRequest {
	return &AlibabaSellerVendorWriteClientAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSellerVendorWriteClientAPIRequest) Reset() {
	r._appId = ""
	r._paramThirdPartyClientDataParams = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSellerVendorWriteClientAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.vendor.write.client"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSellerVendorWriteClientAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSellerVendorWriteClientAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// 开放平台appId
func (r *AlibabaSellerVendorWriteClientAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaSellerVendorWriteClientAPIRequest) GetAppId() string {
	return r._appId
}

// SetParamThirdPartyClientDataParams is ParamThirdPartyClientDataParams Setter
// 要回写的数据
func (r *AlibabaSellerVendorWriteClientAPIRequest) SetParamThirdPartyClientDataParams(_paramThirdPartyClientDataParams *ThirdPartyClientDataParams) error {
	r._paramThirdPartyClientDataParams = _paramThirdPartyClientDataParams
	r.Set("param_third_party_client_data_params", _paramThirdPartyClientDataParams)
	return nil
}

// GetParamThirdPartyClientDataParams ParamThirdPartyClientDataParams Getter
func (r AlibabaSellerVendorWriteClientAPIRequest) GetParamThirdPartyClientDataParams() *ThirdPartyClientDataParams {
	return r._paramThirdPartyClientDataParams
}

var poolAlibabaSellerVendorWriteClientAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSellerVendorWriteClientRequest()
	},
}

// GetAlibabaSellerVendorWriteClientRequest 从 sync.Pool 获取 AlibabaSellerVendorWriteClientAPIRequest
func GetAlibabaSellerVendorWriteClientAPIRequest() *AlibabaSellerVendorWriteClientAPIRequest {
	return poolAlibabaSellerVendorWriteClientAPIRequest.Get().(*AlibabaSellerVendorWriteClientAPIRequest)
}

// ReleaseAlibabaSellerVendorWriteClientAPIRequest 将 AlibabaSellerVendorWriteClientAPIRequest 放入 sync.Pool
func ReleaseAlibabaSellerVendorWriteClientAPIRequest(v *AlibabaSellerVendorWriteClientAPIRequest) {
	v.Reset()
	poolAlibabaSellerVendorWriteClientAPIRequest.Put(v)
}
