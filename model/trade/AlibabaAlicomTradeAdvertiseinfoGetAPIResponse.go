package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomtradeadvertiseinfogetAPIResponse 获取订单上的在信息流投放信息 API返回值
// alibaba.alicom.trade.advertiseinfo.get
//
// 获取订单上的在信息流投放信息
type AlibabaalicomtradeadvertiseinfogetAPIResponse struct {
	model.CommonResponse
	AlibabaalicomtradeadvertiseinfogetAPIResponseModel
}

// AlibabaalicomtradeadvertiseinfogetAPIResponseModel is 获取订单上的在信息流投放信息 成功返回结果
type AlibabaalicomtradeadvertiseinfogetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_trade_advertiseinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 信息投放广告信息
	Module *AdvertiseInfoDto `json:"module,omitempty" xml:"module,omitempty"`
}
