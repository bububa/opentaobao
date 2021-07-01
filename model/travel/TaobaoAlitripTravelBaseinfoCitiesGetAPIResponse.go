package travel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelBaseinfoCitiesGetAPIResponse
【API3.0】度假线路商品发布时基础信息获取接口：地址数据查询 API返回值
taobao.alitrip.travel.baseinfo.cities.get

旅行度假新商品发布时可用的扩展接口，用于获取可用的出发地或目的地城市列表。 */
type TaobaoAlitripTravelBaseinfoCitiesGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelBaseinfoCitiesGetAPIResponseModel
}

// TaobaoAlitripTravelBaseinfoCitiesGetAPIResponseModel is 【API3.0】度假线路商品发布时基础信息获取接口：地址数据查询 成功返回结果
type TaobaoAlitripTravelBaseinfoCitiesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_baseinfo_cities_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 地区级联城市列表，返回数据为json数组结构的字符串
	IocInfos string `json:"ioc_infos,omitempty" xml:"ioc_infos,omitempty"`
}
