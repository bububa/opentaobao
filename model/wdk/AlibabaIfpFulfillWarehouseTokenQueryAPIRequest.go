package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIfpFulfillWarehouseTokenQueryAPIRequest
同城令牌打印接口 API请求
alibaba.ifp.fulfill.warehouse.token.query

用于仓内作业打印包裹信息 */
type AlibabaIfpFulfillWarehouseTokenQueryAPIRequest struct {
	model.Params
	// 入参
	_packageQueryDTO *PackageQueryDto
}

// New
