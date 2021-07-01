package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAscpPoItemQueryAPIRequest
AliExpress采购单明细查询API API请求
aliexpress.ascp.po.item.query

AE 供应链仓发 采购单明细查询 */
type AliexpressAscpPoItemQueryAPIRequest struct {
	model.Params
	// demo
	_purchaseOrderItemQuery *PurchaseOrderItemQueryDto
}

// New
