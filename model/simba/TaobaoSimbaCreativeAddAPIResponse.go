package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativeAddAPIResponse 增加创意 API返回值
// taobao.simba.creative.add
//
// 创建一个创意
type TaobaoSimbaCreativeAddAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCreativeAddAPIResponseModel
}

// TaobaoSimbaCreativeAddAPIResponseModel is 增加创意 成功返回结果
type TaobaoSimbaCreativeAddAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_creative_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 新增加的创意对象
	Creative *Creative `json:"creative,omitempty" xml:"creative,omitempty"`
}
