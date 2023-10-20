package wenyuvideo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YoukuwenyuvideopersionsearchAPIResponse 根据人物名称查询人物列表 API返回值
// youku.wenyuvideo.persion.search
//
// 根据人物名称查询人物列表
type YoukuwenyuvideopersionsearchAPIResponse struct {
	model.CommonResponse
	YoukuwenyuvideopersionsearchAPIResponseModel
}

// YoukuwenyuvideopersionsearchAPIResponseModel is 根据人物名称查询人物列表 成功返回结果
type YoukuwenyuvideopersionsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_wenyuvideo_persion_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *YoukuwenyuvideopersionsearchResult `json:"result,omitempty" xml:"result,omitempty"`
}
