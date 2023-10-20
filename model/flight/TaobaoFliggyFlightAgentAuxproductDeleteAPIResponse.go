package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofliggyflightagentauxproductdeleteAPIResponse 飞猪机票辅营商品删除 API返回值
// taobao.fliggy.flight.agent.auxproduct.delete
//
// 廉航辅营产品删除接口
type TaobaofliggyflightagentauxproductdeleteAPIResponse struct {
	model.CommonResponse
	TaobaofliggyflightagentauxproductdeleteAPIResponseModel
}

// TaobaofliggyflightagentauxproductdeleteAPIResponseModel is 飞猪机票辅营商品删除 成功返回结果
type TaobaofliggyflightagentauxproductdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"fliggy_flight_agent_auxproduct_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DelAuxProductsRs `json:"result,omitempty" xml:"result,omitempty"`
}
