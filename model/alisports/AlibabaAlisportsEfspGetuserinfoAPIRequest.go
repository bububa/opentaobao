package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsEfspGetuserinfoAPIRequest
获取用户详细信息 API请求
alibaba.alisports.efsp.getuserinfo

阿里体育-体育健身-获取用户详细信息 */
type AlibabaAlisportsEfspGetuserinfoAPIRequest struct {
	model.Params
	// 支付宝ID
	_alipayId string
}

// New
