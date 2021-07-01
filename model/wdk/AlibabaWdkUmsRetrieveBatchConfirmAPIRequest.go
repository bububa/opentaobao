package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsRetrieveBatchConfirmAPIRequest
批量消息确认 API请求
alibaba.wdk.ums.retrieve.batch.confirm

批量消息确认 */
type AlibabaWdkUmsRetrieveBatchConfirmAPIRequest struct {
	model.Params
	// warehouse_code
	_warehouseCode string
	// warehouse_code
	_uuids []string
}

// New
