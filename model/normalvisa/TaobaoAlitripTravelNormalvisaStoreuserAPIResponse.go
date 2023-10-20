package normalvisa

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelnormalvisastoreuserAPIResponse 代填办理人信息 API返回值
// taobao.alitrip.travel.normalvisa.storeuser
//
// 卖家代填买家填写办理人信息
type TaobaoalitriptravelnormalvisastoreuserAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelnormalvisastoreuserAPIResponseModel
}

// TaobaoalitriptravelnormalvisastoreuserAPIResponseModel is 代填办理人信息 成功返回结果
type TaobaoalitriptravelnormalvisastoreuserAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_normalvisa_storeuser_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果：包含results数组代表成功
	Result *TaobaoalitriptravelnormalvisastoreuserResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
