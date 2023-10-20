package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacreativedeleteAPIResponse 删除创意 API返回值
// taobao.simba.creative.delete
//
// 删除一个创意
type TaobaosimbacreativedeleteAPIResponse struct {
	model.CommonResponse
	TaobaosimbacreativedeleteAPIResponseModel
}

// TaobaosimbacreativedeleteAPIResponseModel is 删除创意 成功返回结果
type TaobaosimbacreativedeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_creative_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 被删除的创意对象
	Creative *Creative `json:"creative,omitempty" xml:"creative,omitempty"`
}
