package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenReceiverinfoQueryAPIRequest
OAID 收件人信息解密接口 API请求
taobao.qimen.receiverinfo.query

WMS 调用该接口，通过 OAID 查询解密后的收件人信息 */
type TaobaoQimenReceiverinfoQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenReceiverinfoQueryRequest
}

// New
