package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallServiceSettleadjustmentGetAPIResponse
查询结算调整单单条记录 API返回值
tmall.service.settleadjustment.get

提供给服务商通过结算调整单id获取结算调整单信息 */
type TmallServiceSettleadjustmentGetAPIResponse struct {
	model.CommonResponse
	TmallServiceSettleadjustmentGetAPIResponseModel
}

// TmallServiceSettleadjustmentGetAPIResponseModel is 查询结算调整单单条记录 成功返回结果
type TmallServiceSettleadjustmentGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_settleadjustment_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServiceSettleadjustmentGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
