package normalvisa

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelnormalvisagetdetailAPIResponse 获取单笔订单的详情 API返回值
// taobao.alitrip.travel.normalvisa.getdetail
//
// 获取单笔签证的详细记录
type TaobaoalitriptravelnormalvisagetdetailAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelnormalvisagetdetailAPIResponseModel
}

// TaobaoalitriptravelnormalvisagetdetailAPIResponseModel is 获取单笔订单的详情 成功返回结果
type TaobaoalitriptravelnormalvisagetdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_normalvisa_getdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoalitriptravelnormalvisagetdetailResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
