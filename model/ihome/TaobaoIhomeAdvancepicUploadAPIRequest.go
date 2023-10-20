package ihome

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoIhomeAdvancepicUploadAPIRequest ihome图片上传 API请求
// taobao.ihome.advancepic.upload
//
// ihome 定制业务编辑器投稿素材上传
type TaobaoIhomeAdvancepicUploadAPIRequest struct {
	model.Params
	// 图片类
	_materials []AdvancePicMaterialDto
}

// NewTaobaoIhomeAdvancepicUploadRequest 初始化TaobaoIhomeAdvancepicUploadAPIRequest对象
func NewTaobaoIhomeAdvancepicUploadRequest() *TaobaoIhomeAdvancepicUploadAPIRequest {
	return &TaobaoIhomeAdvancepicUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoIhomeAdvancepicUploadAPIRequest) Reset() {
	r._materials = r._materials[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoIhomeAdvancepicUploadAPIRequest) GetApiMethodName() string {
	return "taobao.ihome.advancepic.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoIhomeAdvancepicUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoIhomeAdvancepicUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMaterials is Materials Setter
// 图片类
func (r *TaobaoIhomeAdvancepicUploadAPIRequest) SetMaterials(_materials []AdvancePicMaterialDto) error {
	r._materials = _materials
	r.Set("materials", _materials)
	return nil
}

// GetMaterials Materials Getter
func (r TaobaoIhomeAdvancepicUploadAPIRequest) GetMaterials() []AdvancePicMaterialDto {
	return r._materials
}

var poolTaobaoIhomeAdvancepicUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoIhomeAdvancepicUploadRequest()
	},
}

// GetTaobaoIhomeAdvancepicUploadRequest 从 sync.Pool 获取 TaobaoIhomeAdvancepicUploadAPIRequest
func GetTaobaoIhomeAdvancepicUploadAPIRequest() *TaobaoIhomeAdvancepicUploadAPIRequest {
	return poolTaobaoIhomeAdvancepicUploadAPIRequest.Get().(*TaobaoIhomeAdvancepicUploadAPIRequest)
}

// ReleaseTaobaoIhomeAdvancepicUploadAPIRequest 将 TaobaoIhomeAdvancepicUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoIhomeAdvancepicUploadAPIRequest(v *TaobaoIhomeAdvancepicUploadAPIRequest) {
	v.Reset()
	poolTaobaoIhomeAdvancepicUploadAPIRequest.Put(v)
}
