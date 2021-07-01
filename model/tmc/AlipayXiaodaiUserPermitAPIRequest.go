package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlipayXiaodaiUserPermitAPIRequest
阿里金融用户授权 API请求
alipay.xiaodai.user.permit

阿里金融为用户开通消息通道接口 */
type AlipayXiaodaiUserPermitAPIRequest struct {
	model.Params
	// 用户数字ID
	_userId int64
}

// New
