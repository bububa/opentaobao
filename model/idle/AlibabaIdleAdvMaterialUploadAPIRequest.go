package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAdvMaterialUploadAPIRequest 闲鱼用户增长素材中心素材上传接口 API请求
// alibaba.idle.adv.material.upload
//
// 闲鱼用户增长素材中心素材上传接口
type AlibabaIdleAdvMaterialUploadAPIRequest struct {
	model.Params
	// 素材上传参数
	_uploadTopParam *IdleAdvMaterialUploadTopParam
}

// NewAlibabaIdleAdvMaterialUploadRequest 初始化AlibabaIdleAdvMaterialUploadAPIRequest对象
func NewAlibabaIdleAdvMaterialUploadRequest() *AlibabaIdleAdvMaterialUploadAPIRequest {
	return &AlibabaIdleAdvMaterialUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAdvMaterialUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.adv.material.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAdvMaterialUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleAdvMaterialUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUploadTopParam is UploadTopParam Setter
// 素材上传参数
func (r *AlibabaIdleAdvMaterialUploadAPIRequest) SetUploadTopParam(_uploadTopParam *IdleAdvMaterialUploadTopParam) error {
	r._uploadTopParam = _uploadTopParam
	r.Set("upload_top_param", _uploadTopParam)
	return nil
}

// GetUploadTopParam UploadTopParam Getter
func (r AlibabaIdleAdvMaterialUploadAPIRequest) GetUploadTopParam() *IdleAdvMaterialUploadTopParam {
	return r._uploadTopParam
}
