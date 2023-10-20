package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbownerdeleteAPIResponse 民宿房东删除接口 API返回值
// taobao.xhotel.bnbowner.delete
//
// 民宿房东删除接口，删除房东后，对应的门店及房源会同步删除，请谨慎使用
type TaobaoxhotelbnbownerdeleteAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelbnbownerdeleteAPIResponseModel
}

// TaobaoxhotelbnbownerdeleteAPIResponseModel is 民宿房东删除接口 成功返回结果
type TaobaoxhotelbnbownerdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbowner_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoxhotelbnbownerdeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
