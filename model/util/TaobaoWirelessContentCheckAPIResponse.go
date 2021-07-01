package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWirelessContentCheckAPIResponse
无线开放内容检查 API返回值
taobao.wireless.content.check

无线开放内容检查，提供涉黄暴力政治文本检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70439.html" target="blank">阿里云内容安全</a> */
type TaobaoWirelessContentCheckAPIResponse struct {
	model.CommonResponse
	TaobaoWirelessContentCheckAPIResponseModel
}

// TaobaoWirelessContentCheckAPIResponseModel is 无线开放内容检查 成功返回结果
type TaobaoWirelessContentCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"wireless_content_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 检查结果
	CheckResults []Checkpoints `json:"check_results,omitempty" xml:"check_results>checkpoints,omitempty"`
	// 综合结果建议。建议用户执行的操作，取值范围： pass：文本正常； review：需要人工审核； block：文本违规，可以直接删除或者做限制处理
	Suggestion string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
}
