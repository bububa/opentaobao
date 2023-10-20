package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoDeleteAPIResponse 民宿卖家活动删除 API返回值
// taobao.xhotel.bnbpromo.delete
//
// 民宿删除营销活动
type TaobaoXhotelBnbpromoDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbpromoDeleteAPIResponseModel
}

// TaobaoXhotelBnbpromoDeleteAPIResponseModel is 民宿卖家活动删除 成功返回结果
type TaobaoXhotelBnbpromoDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbpromo_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoXhotelBnbpromoDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
