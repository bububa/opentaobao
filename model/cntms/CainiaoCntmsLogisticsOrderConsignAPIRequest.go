package cntms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCntmsLogisticsOrderConsignAPIRequest
菜鸟配商家仓库发货 API请求
cainiao.cntms.logistics.order.consign

商家包装打印面单结束后，通知菜鸟包裹要发货 */
type CainiaoCntmsLogisticsOrderConsignAPIRequest struct {
	model.Params
	// 配送发货信息
	_content *CnTmsLogisticsOrderConsignContent
}

// New
