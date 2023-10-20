package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordsbyadgroupidgetAPIResponse 取得一个推广组的所有关键词 API返回值
// taobao.simba.keywordsbyadgroupid.get
//
// 取得一个推广组的所有关键词
type TaobaosimbakeywordsbyadgroupidgetAPIResponse struct {
	model.CommonResponse
	TaobaosimbakeywordsbyadgroupidgetAPIResponseModel
}

// TaobaosimbakeywordsbyadgroupidgetAPIResponseModel is 取得一个推广组的所有关键词 成功返回结果
type TaobaosimbakeywordsbyadgroupidgetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywordsbyadgroupid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 取得的关键词列表
	Keywords []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
}
