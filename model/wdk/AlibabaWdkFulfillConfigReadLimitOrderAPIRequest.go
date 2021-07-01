package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillConfigReadLimitOrderAPIRequest
根据仓code查询仓限单配置 API请求
alibaba.wdk.fulfill.config.read.limit.order

根据仓code查询仓限单配置 */
type AlibabaWdkFulfillConfigReadLimitOrderAPIRequest struct {
	model.Params
	// 仓code集合
	_warehouseCodeList []string
}

// New
