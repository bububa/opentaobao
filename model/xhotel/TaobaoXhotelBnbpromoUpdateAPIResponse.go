package xhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbpromoupdateAPIResponse 民宿营销活动更新 API返回值
// taobao.xhotel.bnbpromo.update
//
// 全量更新对应外部活动code相关的营销活动信息
type TaobaoxhotelbnbpromoupdateAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelbnbpromoupdateAPIResponseModel
}

// TaobaoxhotelbnbpromoupdateAPIResponseModel is 民宿营销活动更新 成功返回结果
type TaobaoxhotelbnbpromoupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbpromo_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoxhotelbnbpromoupdateResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
