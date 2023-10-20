package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwayaccountofflinelayeredfindAPIResponse 获取账户历史报表30天转化周期 API返回值
// taobao.subway.account.offline.layeredfind
//
// 获取账户历史报表
type TaobaosubwayaccountofflinelayeredfindAPIResponse struct {
	model.CommonResponse
	TaobaosubwayaccountofflinelayeredfindAPIResponseModel
}

// TaobaosubwayaccountofflinelayeredfindAPIResponseModel is 获取账户历史报表30天转化周期 成功返回结果
type TaobaosubwayaccountofflinelayeredfindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_account_offline_layeredfind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result []ReportResultTopDto `json:"result,omitempty" xml:"result>report_result_top_dto,omitempty"`
	// 提示
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
