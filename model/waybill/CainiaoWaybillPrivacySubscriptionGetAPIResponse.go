package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybillprivacysubscriptiongetAPIResponse 隐私面单商家订购查询 API返回值
// cainiao.waybill.privacy.subscription.get
//
// ISV查询商家是否订购隐私面单
type CainiaowaybillprivacysubscriptiongetAPIResponse struct {
	model.CommonResponse
	CainiaowaybillprivacysubscriptiongetAPIResponseModel
}

// CainiaowaybillprivacysubscriptiongetAPIResponseModel is 隐私面单商家订购查询 成功返回结果
type CainiaowaybillprivacysubscriptiongetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_privacy_subscription_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *CainiaowaybillprivacysubscriptiongetResult `json:"result,omitempty" xml:"result,omitempty"`
}
