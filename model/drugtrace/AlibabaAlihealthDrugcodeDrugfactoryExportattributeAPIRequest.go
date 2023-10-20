package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodedrugfactoryexportattributeAPIRequest 导出所有项目的药物属性和药品信息 API请求
// alibaba.alihealth.drugcode.drugfactory.exportattribute
//
// 导出所有项目的药物属性和药品信息
type AlibabaalihealthdrugcodedrugfactoryexportattributeAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
}

// NewAlibabaalihealthdrugcodedrugfactoryexportattributeRequest 初始化AlibabaalihealthdrugcodedrugfactoryexportattributeAPIRequest对象
func NewAlibabaalihealthdrugcodedrugfactoryexportattributeRequest() *AlibabaalihealthdrugcodedrugfactoryexportattributeAPIRequest {
	return &AlibabaalihealthdrugcodedrugfactoryexportattributeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodedrugfactoryexportattributeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.exportattribute"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodedrugfactoryexportattributeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodedrugfactoryexportattributeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugcodedrugfactoryexportattributeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodedrugfactoryexportattributeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
