package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopensellerbizlogistictimeruleAPIResponse 商家自定义发货时效 API返回值
// taobao.open.seller.biz.logistic.time.rule
//
// 服务商回传商家自定义发货时效
type TaobaoopensellerbizlogistictimeruleAPIResponse struct {
	model.CommonResponse
	TaobaoopensellerbizlogistictimeruleAPIResponseModel
}

// TaobaoopensellerbizlogistictimeruleAPIResponseModel is 商家自定义发货时效 成功返回结果
type TaobaoopensellerbizlogistictimeruleAPIResponseModel struct {
	XMLName xml.Name `xml:"open_seller_biz_logistic_time_rule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
