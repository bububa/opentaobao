package medicalbase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest
号源直连订单状态同步接口 API请求
alibaba.alihealth.medicalbase.order.status.sync

互联网医院isv批量通过接口批量导入 */
type AlibabaAlihealthMedicalbaseOrderStatusSyncAPIRequest struct {
	model.Params
	// 订单信息
	_orderlSyncDTO *OrderlSyncDto
}

// New
