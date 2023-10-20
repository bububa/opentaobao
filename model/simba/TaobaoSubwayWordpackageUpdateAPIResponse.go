package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayWordpackageUpdateAPIResponse 批量更新词包 API返回值
// taobao.subway.wordpackage.update
//
// 批量更新词包
type TaobaoSubwayWordpackageUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayWordpackageUpdateAPIResponseModel
}

// TaobaoSubwayWordpackageUpdateAPIResponseModel is 批量更新词包 成功返回结果
type TaobaoSubwayWordpackageUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_wordpackage_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoSubwayWordpackageUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
