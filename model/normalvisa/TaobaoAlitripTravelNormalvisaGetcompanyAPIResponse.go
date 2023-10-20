package normalvisa

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelnormalvisagetcompanyAPIResponse 获取物流公司信息 API返回值
// taobao.alitrip.travel.normalvisa.getcompany
//
// 获取物流公司信息
type TaobaoalitriptravelnormalvisagetcompanyAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelnormalvisagetcompanyAPIResponseModel
}

// TaobaoalitriptravelnormalvisagetcompanyAPIResponseModel is 获取物流公司信息 成功返回结果
type TaobaoalitriptravelnormalvisagetcompanyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_normalvisa_getcompany_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果，有返回代表成功
	Result *TaobaoalitriptravelnormalvisagetcompanyResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
