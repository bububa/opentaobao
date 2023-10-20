package util

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPictureQnaigcUploadAPIRequest) Reset() {
	r._uploadRequest = nil
	r.Params.ToZero()
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

var poolTaobaoPictureQnaigcUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPictureQnaigcUploadRequest()
	},
}

// GetTaobaoPictureQnaigcUploadRequest 从 sync.Pool 获取 TaobaoPictureQnaigcUploadAPIRequest
func GetTaobaoPictureQnaigcUploadAPIRequest() *TaobaoPictureQnaigcUploadAPIRequest {
	return poolTaobaoPictureQnaigcUploadAPIRequest.Get().(*TaobaoPictureQnaigcUploadAPIRequest)
}

// ReleaseTaobaoPictureQnaigcUploadAPIRequest 将 TaobaoPictureQnaigcUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoPictureQnaigcUploadAPIRequest(v *TaobaoPictureQnaigcUploadAPIRequest) {
	v.Reset()
	poolTaobaoPictureQnaigcUploadAPIRequest.Put(v)
}
