package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelCommoninvoiceRemoveAPIRequest
常用发票信息删除接口 API请求
taobao.xhotel.commoninvoice.remove

常用发票信息删除接口 */
type TaobaoXhotelCommoninvoiceRemoveAPIRequest struct {
	model.Params
	// 发票id
	_invoiceId int64
	// 用户名
	_userNick string
}

// New
