package alihealthcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse
alibaba.alihealth.baby.baseinfo.order.sync API返回值
alibaba.alihealth.baby.baseinfo.order.sync

育学园将订单信息回传给我们 */
type AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponseModel
}

// AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponseModel is alibaba.alihealth.baby.baseinfo.order.sync 成功返回结果
type AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_baby_baseinfo_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作提示
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 操作码
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
}
