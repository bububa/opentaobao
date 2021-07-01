package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkDgNewuserOrderSumAPIResponse
淘宝客-推广者-拉新活动对应数据查询 API返回值
taobao.tbk.dg.newuser.order.sum

拉新活动汇总API */
type TaobaoTbkDgNewuserOrderSumAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgNewuserOrderSumAPIResponseModel
}

// TaobaoTbkDgNewuserOrderSumAPIResponseModel is 淘宝客-推广者-拉新活动对应数据查询 成功返回结果
type TaobaoTbkDgNewuserOrderSumAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_newuser_order_sum_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Results *TaobaoTbkDgNewuserOrderSumData `json:"results,omitempty" xml:"results,omitempty"`
}
