package tmallhk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraceplatformTicketPictureUploadAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.ticket.picture.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraceplatformTicketPictureUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizOrderId Setter
// 子订单号
func (r *TmallTraceplatformTicketPictureUploadAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// Get BizOrderId Getter
func (r TmallTraceplatformTicketPictureUploadAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// Set is File Setter
// 图片二进制流，只支持jpg/jpeg/png格式
func (r *TmallTraceplatformTicketPictureUploadAPIRequest) SetFile(_file *model.File) error {
	r._file = _file
	r.Set("file", _file)
	return nil
}

// Get File Getter
func (r TmallTraceplatformTicketPictureUploadAPIRequest) GetFile() *model.File {
	return r._file
}
