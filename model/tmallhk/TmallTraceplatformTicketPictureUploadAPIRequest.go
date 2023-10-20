package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltraceplatformticketpictureuploadAPIRequest 上传小票图片 API请求
// tmall.traceplatform.ticket.picture.upload
//
// uploadPicture
type TmalltraceplatformticketpictureuploadAPIRequest struct {
	model.Params
	// 子订单号
	_bizOrderId int64
	// 图片二进制流，只支持jpg/jpeg/png格式
	_file *model.File
}

// NewTmalltraceplatformticketpictureuploadRequest 初始化TmalltraceplatformticketpictureuploadAPIRequest对象
func NewTmalltraceplatformticketpictureuploadRequest() *TmalltraceplatformticketpictureuploadAPIRequest {
	return &TmalltraceplatformticketpictureuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltraceplatformticketpictureuploadAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.ticket.picture.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltraceplatformticketpictureuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltraceplatformticketpictureuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 子订单号
func (r *TmalltraceplatformticketpictureuploadAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmalltraceplatformticketpictureuploadAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// SetFile is File Setter
// 图片二进制流，只支持jpg/jpeg/png格式
func (r *TmalltraceplatformticketpictureuploadAPIRequest) SetFile(_file *model.File) error {
	r._file = _file
	r.Set("file", _file)
	return nil
}

// GetFile File Getter
func (r TmalltraceplatformticketpictureuploadAPIRequest) GetFile() *model.File {
	return r._file
}
