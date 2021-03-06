package wangwang

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWangwangAbstractLogqueryAPIResponse 模糊聊天记录查询 API返回值
// taobao.wangwang.abstract.logquery
//
// 模糊聊天记录查询
type TaobaoWangwangAbstractLogqueryAPIResponse struct {
	model.CommonResponse
	TaobaoWangwangAbstractLogqueryAPIResponseModel
}

// TaobaoWangwangAbstractLogqueryAPIResponseModel is 模糊聊天记录查询 成功返回结果
type TaobaoWangwangAbstractLogqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wangwang_abstract_logquery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 当direction为1时有效，url列表
	UrlLists []UrlList `json:"url_lists,omitempty" xml:"url_lists>url_list,omitempty"`
	// 当direction为1时有效，关键词统计列表
	WordCountLists []WordCountList `json:"word_count_lists,omitempty" xml:"word_count_lists>word_count_list,omitempty"`
	// 消息列表
	MsgLists []MsgList `json:"msg_lists,omitempty" xml:"msg_lists>msg_list,omitempty"`
	// 例如单词长度太长等。<br/>当ret_code不为0时这个值存在。
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 卖家id
	FromId string `json:"from_id,omitempty" xml:"from_id,omitempty"`
	// 买家id
	ToId string `json:"to_id,omitempty" xml:"to_id,omitempty"`
	// 当is_sliced为1时有效
	NextKey string `json:"next_key,omitempty" xml:"next_key,omitempty"`
	// 0或-1，表示错误或正确，错误时有错误信息.<br/>为-1时就只有error_msg字段是有效的。
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 0或1<br/>其他返回时，是由于用户名等参数设置有误等引起的远端服务错误
	IsSliced int64 `json:"is_sliced,omitempty" xml:"is_sliced,omitempty"`
}
