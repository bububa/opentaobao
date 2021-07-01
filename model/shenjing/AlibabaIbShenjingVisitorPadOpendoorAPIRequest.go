package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIbShenjingVisitorPadOpendoorAPIRequest
访客发起开门 API请求
alibaba.ib.shenjing.visitor.pad.opendoor

访客PAD端录入完人脸后，可以点击开门按钮开门。 */
type AlibabaIbShenjingVisitorPadOpendoorAPIRequest struct {
	model.Params
	// 访客标识
	_id string
	// padid
	_padId string
}

// New
