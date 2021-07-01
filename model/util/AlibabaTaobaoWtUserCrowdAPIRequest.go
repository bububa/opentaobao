package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTaobaoWtUserCrowdAPIRequest
校验用户是否为新人 API请求
alibaba.taobao.wt.user.crowd

增加isv接口
根据入参手机号判断是否为新人类型 */
type AlibabaTaobaoWtUserCrowdAPIRequest struct {
	model.Params
	// 手机号
	_phone int64
}

// New
