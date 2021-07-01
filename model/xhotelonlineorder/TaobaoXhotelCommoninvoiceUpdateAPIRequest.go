package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelCommoninvoiceUpdateAPIRequest
常用发票信息更新接口 API请求
taobao.xhotel.commoninvoice.update

常用发票信息更新接口(根据用户id,发票抬头和发票属性或发票id进行更新,没有则添加) */
type TaobaoXhotelCommoninvoiceUpdateAPIRequest struct {
	model.Params
	// 无
	_commonInvoiceInfoParam *CommonInvoiceInfo
}

// New
