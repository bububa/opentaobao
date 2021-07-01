package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmShopListGetAPIRequest
获取商圈商户信息列表 API请求
alibaba.westcrm.shop.list.get

获取商圈商户信息列表 */
type AlibabaWestcrmShopListGetAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
}

// New
