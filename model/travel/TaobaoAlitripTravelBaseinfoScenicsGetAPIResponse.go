package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse 【API3.0】基础信息获取接口：景点数据查询 API返回值
// taobao.alitrip.travel.baseinfo.scenics.get
//
// 商品发布辅助接口，用于飞猪度假或门票商品发布时 获取可用的景点（及景点下收费项目）信息列表。
type TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelBaseinfoScenicsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelBaseinfoScenicsGetAPIResponseModel).Reset()
}

// TaobaoAlitripTravelBaseinfoScenicsGetAPIResponseModel is 【API3.0】基础信息获取接口：景点数据查询 成功返回结果
type TaobaoAlitripTravelBaseinfoScenicsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_baseinfo_scenics_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回详细景点信息，返回数据为json数组结构的字符串
	ScenicInfos string `json:"scenic_infos,omitempty" xml:"scenic_infos,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelBaseinfoScenicsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ScenicInfos = ""
}

var poolTaobaoAlitripTravelBaseinfoScenicsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse)
	},
}

// GetTaobaoAlitripTravelBaseinfoScenicsGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse
func GetTaobaoAlitripTravelBaseinfoScenicsGetAPIResponse() *TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse {
	return poolTaobaoAlitripTravelBaseinfoScenicsGetAPIResponse.Get().(*TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse)
}

// ReleaseTaobaoAlitripTravelBaseinfoScenicsGetAPIResponse 将 TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelBaseinfoScenicsGetAPIResponse(v *TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelBaseinfoScenicsGetAPIResponse.Put(v)
}
