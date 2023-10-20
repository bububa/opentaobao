package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodedrugfactoryexportprojectAPIRequest 导出项目和药品目录 API请求
// alibaba.alihealth.drugcode.drugfactory.exportproject
//
// 导出临床项目及其药品目录
type AlibabaalihealthdrugcodedrugfactoryexportprojectAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
}

// NewAlibabaalihealthdrugcodedrugfactoryexportprojectRequest 初始化AlibabaalihealthdrugcodedrugfactoryexportprojectAPIRequest对象
func NewAlibabaalihealthdrugcodedrugfactoryexportprojectRequest() *AlibabaalihealthdrugcodedrugfactoryexportprojectAPIRequest {
	return &AlibabaalihealthdrugcodedrugfactoryexportprojectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodedrugfactoryexportprojectAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.exportproject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodedrugfactoryexportprojectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodedrugfactoryexportprojectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugcodedrugfactoryexportprojectAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodedrugfactoryexportprojectAPIRequest) GetRefEntId() string {
	return r._refEntId
}
