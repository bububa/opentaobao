package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTraceplatformTicketPictureUploadAPIRequest
上传小票图片 API请求
tmall.traceplatform.ticket.picture.upload

uploadPicture */
type TmallTraceplatformTicketPictureUploadAPIRequest struct {
	model.Params
	// 子订单号
	_bizOrderId int64
	// 图片二进制流，只支持jpg/jpeg/png格式
	_file *model.File
}

// New
