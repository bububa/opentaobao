package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIRequest 获取加密公钥 API请求
// alibaba.alihealth.drugcode.drugfactory.getencrptypk
//
// 获取服务端给药厂用来加密的公钥
type AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIRequest struct {
	model.Params
	// 企业Id
	_refEntId string
}

// NewAlibabaalihealthdrugcodedrugfactorygetencrptypkRequest 初始化AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIRequest对象
func NewAlibabaalihealthdrugcodedrugfactorygetencrptypkRequest() *AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIRequest {
	return &AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.getencrptypk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业Id
func (r *AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIRequest) GetRefEntId() string {
	return r._refEntId
}
