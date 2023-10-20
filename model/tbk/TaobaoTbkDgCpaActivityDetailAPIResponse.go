package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgcpaactivitydetailAPIResponse 淘宝客-推广者-CPA活动执行明细 API返回值
// taobao.tbk.dg.cpa.activity.detail
//
// 淘宝客获取CPA活动具体执行效果的明细数据（含预估和结算维度）
type TaobaotbkdgcpaactivitydetailAPIResponse struct {
	model.CommonResponse
	TaobaotbkdgcpaactivitydetailAPIResponseModel
}

// TaobaotbkdgcpaactivitydetailAPIResponseModel is 淘宝客-推广者-CPA活动执行明细 成功返回结果
type TaobaotbkdgcpaactivitydetailAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_cpa_activity_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaotbkdgcpaactivitydetailResult `json:"result,omitempty" xml:"result,omitempty"`
}
