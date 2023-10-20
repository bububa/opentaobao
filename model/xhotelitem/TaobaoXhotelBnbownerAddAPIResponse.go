package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbowneraddAPIResponse 民宿房东信息添加 API返回值
// taobao.xhotel.bnbowner.add
//
// 添加和更新民宿房东的信息
type TaobaoxhotelbnbowneraddAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelbnbowneraddAPIResponseModel
}

// TaobaoxhotelbnbowneraddAPIResponseModel is 民宿房东信息添加 成功返回结果
type TaobaoxhotelbnbowneraddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbowner_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoxhotelbnbowneraddResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
