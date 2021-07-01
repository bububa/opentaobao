package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalItemStatusAPIRequest
商品上下架 API请求
alibaba.alihealth.medical.item.status

医生通三方机构平台进行服务商品上下架操作 */
type AlibabaAlihealthMedicalItemStatusAPIRequest struct {
	model.Params
	// 请求入参
	_shelfrequest *ThirdAgencyUpDownShelfRequest
}

// New
