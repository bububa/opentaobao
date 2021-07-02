package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiIdGetAPIResponse 获取请求id API返回值
// alibaba.happytrip.taxi.id.get
//
// 获取订单号
type AlibabaHappytripTaxiIdGetAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiIdGetAPIResponseModel
}

// AlibabaHappytripTaxiIdGetAPIResponseModel is 获取请求id 成功返回结果
type AlibabaHappytripTaxiIdGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_id_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// id获取结果
	Data *GetIdResult `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}
