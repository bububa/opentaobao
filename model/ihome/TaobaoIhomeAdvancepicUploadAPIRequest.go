package ihome

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoihomeadvancepicuploadAPIRequest ihome图片上传 API请求
// taobao.ihome.advancepic.upload
//
// ihome 定制业务编辑器投稿素材上传
type TaobaoihomeadvancepicuploadAPIRequest struct {
	model.Params
	// 图片类
	_materials []AdvancePicMaterialDto
}

// NewTaobaoihomeadvancepicuploadRequest 初始化TaobaoihomeadvancepicuploadAPIRequest对象
func NewTaobaoihomeadvancepicuploadRequest() *TaobaoihomeadvancepicuploadAPIRequest {
	return &TaobaoihomeadvancepicuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoihomeadvancepicuploadAPIRequest) GetApiMethodName() string {
	return "taobao.ihome.advancepic.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoihomeadvancepicuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoihomeadvancepicuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterials is Materials Setter
// 图片类
func (r *TaobaoihomeadvancepicuploadAPIRequest) SetMaterials(_materials []AdvancePicMaterialDto) error {
	r._materials = _materials
	r.Set("materials", _materials)
	return nil
}

// GetMaterials Materials Getter
func (r TaobaoihomeadvancepicuploadAPIRequest) GetMaterials() []AdvancePicMaterialDto {
	return r._materials
}
