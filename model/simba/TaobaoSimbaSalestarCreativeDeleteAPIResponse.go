package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestarcreativedeleteAPIResponse (新)销量明星删除创意相关接口 API返回值
// taobao.simba.salestar.creative.delete
//
// 删除一个创意
type TaobaosimbasalestarcreativedeleteAPIResponse struct {
	model.CommonResponse
	TaobaosimbasalestarcreativedeleteAPIResponseModel
}

// TaobaosimbasalestarcreativedeleteAPIResponseModel is (新)销量明星删除创意相关接口 成功返回结果
type TaobaosimbasalestarcreativedeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_creative_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 被删除的创意对象
	Creative *Creative `json:"creative,omitempty" xml:"creative,omitempty"`
}
