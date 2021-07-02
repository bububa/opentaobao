package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowAccountRptdailylistAPIResponse 获取广告主分日数据 API返回值
// taobao.feedflow.account.rptdailylist
//
// 获取广告主分日数据
type TaobaoFeedflowAccountRptdailylistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowAccountRptdailylistAPIResponseModel
}

// TaobaoFeedflowAccountRptdailylistAPIResponseModel is 获取广告主分日数据 成功返回结果
type TaobaoFeedflowAccountRptdailylistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_account_rptdailylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoFeedflowAccountRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
