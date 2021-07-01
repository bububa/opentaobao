package alihealthcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlipaypfmOrderSyncAPIResponse
订单数据回传接口 API返回值
alibaba.alihealth.alipaypfm.order.sync

订单数据回传接口，各个isv通过我们渠道产生订单需要回传进行统计 */
type AlibabaAlihealthAlipaypfmOrderSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthAlipaypfmOrderSyncAPIResponseModel
}

// AlibabaAlihealthAlipaypfmOrderSyncAPIResponseModel is 订单数据回传接口 成功返回结果
type AlibabaAlihealthAlipaypfmOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_alipaypfm_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
