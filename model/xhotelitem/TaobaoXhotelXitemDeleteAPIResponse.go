package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelXitemDeleteAPIResponse 删除 x 元素 API返回值
// taobao.xhotel.xitem.delete
//
// 删除 x 元素
type TaobaoXhotelXitemDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelXitemDeleteAPIResponseModel
}

// TaobaoXhotelXitemDeleteAPIResponseModel is 删除 x 元素 成功返回结果
type TaobaoXhotelXitemDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_xitem_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoXhotelXitemDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
