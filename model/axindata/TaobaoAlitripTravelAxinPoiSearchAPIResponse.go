package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinpoisearchAPIResponse 景点poi搜索-阿信 API返回值
// taobao.alitrip.travel.axin.poi.search
//
// 给阿信提供景点poi搜索
type TaobaoalitriptravelaxinpoisearchAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinpoisearchAPIResponseModel
}

// TaobaoalitriptravelaxinpoisearchAPIResponseModel is 景点poi搜索-阿信 成功返回结果
type TaobaoalitriptravelaxinpoisearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_poi_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoalitriptravelaxinpoisearchResult `json:"result,omitempty" xml:"result,omitempty"`
}
