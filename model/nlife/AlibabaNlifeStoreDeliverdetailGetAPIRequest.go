package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeStoreDeliverdetailGetAPIRequest
查询发货单详情 API请求
alibaba.nlife.store.deliverdetail.get

查询发货单详情 */
type AlibabaNlifeStoreDeliverdetailGetAPIRequest struct {
	model.Params
	// 发货单号
	_consignNo string
	// 门店id
	_storeId int64
}

// New
