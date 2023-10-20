package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenuidgetbytradeAPIResponse 通过订单获取对应买家的openUID API返回值
// taobao.openuid.get.bytrade
//
// 通过订单获取对应买家的openUID,需要卖家授权
type TaobaoopenuidgetbytradeAPIResponse struct {
	model.CommonResponse
	TaobaoopenuidgetbytradeAPIResponseModel
}

// TaobaoopenuidgetbytradeAPIResponseModel is 通过订单获取对应买家的openUID 成功返回结果
type TaobaoopenuidgetbytradeAPIResponseModel struct {
	XMLName xml.Name `xml:"openuid_get_bytrade_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 当前交易tid对应买家的openuid
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
}
