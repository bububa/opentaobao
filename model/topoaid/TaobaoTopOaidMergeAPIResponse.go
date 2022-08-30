package topoaid

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOaidMergeAPIResponse OAID订单合并 API返回值
// taobao.top.oaid.merge
//
// 基于OAID（收件人ID， Open Addressee ID)做订单合并，确保相同收件人信息的订单合并到相同组。
type TaobaoTopOaidMergeAPIResponse struct {
	model.CommonResponse
	TaobaoTopOaidMergeAPIResponseModel
}

// TaobaoTopOaidMergeAPIResponseModel is OAID订单合并 成功返回结果
type TaobaoTopOaidMergeAPIResponseModel struct {
	XMLName xml.Name `xml:"top_oaid_merge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 合单结果。注意：一个字符串中的多笔订单可以合单，不同字符串间的订单不可合单！！！<br/> 比如，["1111,2222"]表示订单1111和订单2222可合单，["1111","2222"]表示订单1111和订单2222不可以合单。
	TidList []string `json:"tid_list,omitempty" xml:"tid_list>string,omitempty"`
}
