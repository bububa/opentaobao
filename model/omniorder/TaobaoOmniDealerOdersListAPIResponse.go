package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomnidealeroderslistAPIResponse 全渠道经销商订单列表 API返回值
// taobao.omni.dealer.oders.list
//
// 全渠道经销商订单列表查询
type TaobaoomnidealeroderslistAPIResponse struct {
	model.CommonResponse
	TaobaoomnidealeroderslistAPIResponseModel
}

// TaobaoomnidealeroderslistAPIResponseModel is 全渠道经销商订单列表 成功返回结果
type TaobaoomnidealeroderslistAPIResponseModel struct {
	XMLName xml.Name `xml:"omni_dealer_oders_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页结果
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
