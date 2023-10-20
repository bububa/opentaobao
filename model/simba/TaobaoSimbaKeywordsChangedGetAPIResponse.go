package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordschangedgetAPIResponse 分页获取修改过的关键词ID、宝贝id、修改时间 API返回值
// taobao.simba.keywords.changed.get
//
// 分页获取修改过的关键词ID、宝贝id、修改时间
type TaobaosimbakeywordschangedgetAPIResponse struct {
	model.CommonResponse
	TaobaosimbakeywordschangedgetAPIResponseModel
}

// TaobaosimbakeywordschangedgetAPIResponseModel is 分页获取修改过的关键词ID、宝贝id、修改时间 成功返回结果
type TaobaosimbakeywordschangedgetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywords_changed_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词分页对象
	Keywords *KeywordPage `json:"keywords,omitempty" xml:"keywords,omitempty"`
}
