package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelDeleteAPIResponse
删除酒店接口 API返回值
taobao.xhotel.delete

删除飞猪酒店数据接口 */
type TaobaoXhotelDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelDeleteAPIResponseModel
}

// TaobaoXhotelDeleteAPIResponseModel is 删除酒店接口 成功返回结果
type TaobaoXhotelDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除结果
	Result *TaobaoXhotelDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
