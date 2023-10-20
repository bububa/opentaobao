package tmallhk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformTicketPictureUploadAPIRequest 上传小票图片 API请求
// tmall.traceplatform.ticket.picture.upload
//
// uploadPicture
type TmallTraceplatformTicketPictureUploadAPIRequest struct {
	model.Params
	// 子订单号
	_bizOrderId int64
	// 图片二进制流，只支持jpg/jpeg/png格式
	_file *model.File
}

// NewTmallTraceplatformTicketPictureUploadRequest 初始化TmallTraceplatformTicketPictureUploadAPIRequest对象
func NewTmallTraceplatformTicketPictureUploadRequest() *TmallTraceplatformTicketPictureUploadAPIRequest {
	return &TmallTraceplatformTicketPictureUploadAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTraceplatformTicketPictureUploadAPIRequest) Reset() {
	r._bizOrderId = 0
	r._file = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraceplatformTicketPictureUploadAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.ticket.picture.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraceplatformTicketPictureUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTraceplatformTicketPictureUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 子订单号
func (r *TmallTraceplatformTicketPictureUploadAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallTraceplatformTicketPictureUploadAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// SetFile is File Setter
// 图片二进制流，只支持jpg/jpeg/png格式
func (r *TmallTraceplatformTicketPictureUploadAPIRequest) SetFile(_file *model.File) error {
	r._file = _file
	r.Set("file", _file)
	return nil
}

// GetFile File Getter
func (r TmallTraceplatformTicketPictureUploadAPIRequest) GetFile() *model.File {
	return r._file
}

var poolTmallTraceplatformTicketPictureUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTraceplatformTicketPictureUploadRequest()
	},
}

// GetTmallTraceplatformTicketPictureUploadRequest 从 sync.Pool 获取 TmallTraceplatformTicketPictureUploadAPIRequest
func GetTmallTraceplatformTicketPictureUploadAPIRequest() *TmallTraceplatformTicketPictureUploadAPIRequest {
	return poolTmallTraceplatformTicketPictureUploadAPIRequest.Get().(*TmallTraceplatformTicketPictureUploadAPIRequest)
}

// ReleaseTmallTraceplatformTicketPictureUploadAPIRequest 将 TmallTraceplatformTicketPictureUploadAPIRequest 放入 sync.Pool
func ReleaseTmallTraceplatformTicketPictureUploadAPIRequest(v *TmallTraceplatformTicketPictureUploadAPIRequest) {
	v.Reset()
	poolTmallTraceplatformTicketPictureUploadAPIRequest.Put(v)
}
