package koubeimall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonMallNearListAPIResponse 根据POI查询附近商圈列表信息 API返回值
// taobao.koubei.mall.common.mall.near.list
//
// 通过用户/终端设备地理位置POI信息，查询附近商圈信息
type TaobaoKoubeiMallCommonMallNearListAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiMallCommonMallNearListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonMallNearListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiMallCommonMallNearListAPIResponseModel).Reset()
}

// TaobaoKoubeiMallCommonMallNearListAPIResponseModel is 根据POI查询附近商圈列表信息 成功返回结果
type TaobaoKoubeiMallCommonMallNearListAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_mall_near_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoKoubeiMallCommonMallNearListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonMallNearListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiMallCommonMallNearListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonMallNearListAPIResponse)
	},
}

// GetTaobaoKoubeiMallCommonMallNearListAPIResponse 从 sync.Pool 获取 TaobaoKoubeiMallCommonMallNearListAPIResponse
func GetTaobaoKoubeiMallCommonMallNearListAPIResponse() *TaobaoKoubeiMallCommonMallNearListAPIResponse {
	return poolTaobaoKoubeiMallCommonMallNearListAPIResponse.Get().(*TaobaoKoubeiMallCommonMallNearListAPIResponse)
}

// ReleaseTaobaoKoubeiMallCommonMallNearListAPIResponse 将 TaobaoKoubeiMallCommonMallNearListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiMallCommonMallNearListAPIResponse(v *TaobaoKoubeiMallCommonMallNearListAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiMallCommonMallNearListAPIResponse.Put(v)
}
