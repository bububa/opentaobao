package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliicancelAPIResponse 商家取消获取的电子面单号 API返回值
// cainiao.waybill.ii.cancel
//
// 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type CainiaowaybilliicancelAPIResponse struct {
	model.CommonResponse
	CainiaowaybilliicancelAPIResponseModel
}

// CainiaowaybilliicancelAPIResponseModel is 商家取消获取的电子面单号 成功返回结果
type CainiaowaybilliicancelAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用取消是否成功
	CancelResult bool `json:"cancel_result,omitempty" xml:"cancel_result,omitempty"`
}
