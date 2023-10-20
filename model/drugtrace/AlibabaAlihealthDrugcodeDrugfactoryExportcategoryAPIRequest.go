package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIRequest 导出临床药品目录 API请求
// alibaba.alihealth.drugcode.drugfactory.exportcategory
//
// 导出临床药品目录
type AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
}

// NewAlibabaalihealthdrugcodedrugfactoryexportcategoryRequest 初始化AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIRequest对象
func NewAlibabaalihealthdrugcodedrugfactoryexportcategoryRequest() *AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIRequest {
	return &AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.exportcategory"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIRequest) GetRefEntId() string {
	return r._refEntId
}
