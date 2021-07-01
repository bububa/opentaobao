package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSyncedorderQueryAPIRequest
五道口查询同步订单 API请求
alibaba.wdk.syncedorder.query

外部商户查询同步到五道口的订单 */
type AlibabaWdkSyncedorderQueryAPIRequest struct {
	model.Params
	// 门店ID
	_storeId string
	// 序列号
	_serialNum string
}

// New
