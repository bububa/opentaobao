package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAscpPoQueryAPIRequest
AliExpress采购单查询API API请求
aliexpress.ascp.po.query

AE仓发业务采购单查询 */
type AliexpressAscpPoQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_purchaseOrderQuery *PurchaseOrderQueryDto
}

// New
