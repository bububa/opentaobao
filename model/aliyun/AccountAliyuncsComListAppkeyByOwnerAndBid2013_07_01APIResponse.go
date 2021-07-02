package aliyun

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIResponse 根据渠道id和状态列出appkey API返回值
// account.aliyuncs.com.ListAppkeyByOwnerAndBid.2013-07-01
//
// 根据渠道id和状态列出appkey
type AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIResponse struct {
	model.CommonResponse
	AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIResponseModel
}

// AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIResponseModel is 根据渠道id和状态列出appkey 成功返回结果
type AccountAliyuncsComListAppkeyByOwnerAndBid2013_07_01APIResponseModel struct {
	XMLName xml.Name `xml:"account_aliyuncs_com_ListAppkeyByOwnerAndBid_2013-07-01_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return result
	ResultData string `json:"result_data,omitempty" xml:"result_data,omitempty"`
}
