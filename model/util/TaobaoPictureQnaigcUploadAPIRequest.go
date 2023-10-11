package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureQnaigcUploadAPIRequest qnaigc业务线图片上传 API请求
// taobao.picture.qnaigc.upload
//
// qnaigc业务线图片上传
type TaobaoPictureQnaigcUploadAPIRequest struct {
	model.Params
	// qnaigc业务线图片上传请求
	_uploadRequest *UploadRequest
}

// NewTaobaoPictureQnaigcUploadRequest 初始化TaobaoPictureQnaigcUploadAPIRequest对象
func NewTaobaoPictureQnaigcUploadRequest() *TaobaoPictureQnaigcUploadAPIRequest {
	return &TaobaoPictureQnaigcUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureQnaigcUploadAPIRequest) GetApiMethodName() string {
	return "taobao.picture.qnaigc.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureQnaigcUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPictureQnaigcUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUploadRequest is UploadRequest Setter
// qnaigc业务线图片上传请求
func (r *TaobaoPictureQnaigcUploadAPIRequest) SetUploadRequest(_uploadRequest *UploadRequest) error {
	r._uploadRequest = _uploadRequest
	r.Set("upload_request", _uploadRequest)
	return nil
}

// GetUploadRequest UploadRequest Getter
func (r TaobaoPictureQnaigcUploadAPIRequest) GetUploadRequest() *UploadRequest {
	return r._uploadRequest
}
