package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest
pad获取二维码 API请求
alibaba.ib.shenjing.visitor.pad.getqrcodelink

pad获取二维码链接。扫码录入人脸。 */
type AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest struct {
	model.Params
	// 终端id
	_termId string
}

// New
