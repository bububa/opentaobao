package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportquerynotitempromotionAPIResponse 其他主体报表查询 API返回值
// taobao.universalbp.report.query.not.item.promotion
//
// 其他主体报表查询
type TaobaouniversalbpreportquerynotitempromotionAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpreportquerynotitempromotionAPIResponseModel
}

// TaobaouniversalbpreportquerynotitempromotionAPIResponseModel is 其他主体报表查询 成功返回结果
type TaobaouniversalbpreportquerynotitempromotionAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_not_item_promotion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpreportquerynotitempromotionTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
