package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinpoidetailqueryAPIResponse 景点poi详情查询-阿信 API返回值
// taobao.alitrip.travel.axin.poi.detail.query
//
// 景点poi详情查询-阿信
type TaobaoalitriptravelaxinpoidetailqueryAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelaxinpoidetailqueryAPIResponseModel
}

// TaobaoalitriptravelaxinpoidetailqueryAPIResponseModel is 景点poi详情查询-阿信 成功返回结果
type TaobaoalitriptravelaxinpoidetailqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_poi_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoalitriptravelaxinpoidetailqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
