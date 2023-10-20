package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScMembergroupOptionalAPIResponse 工具服务商member组查询、新增接口 API返回值
// taobao.tbk.sc.membergroup.optional
//
// 工具服务商member组查询、新增接口
type TaobaoTbkScMembergroupOptionalAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScMembergroupOptionalAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScMembergroupOptionalAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScMembergroupOptionalAPIResponseModel).Reset()
}

// TaobaoTbkScMembergroupOptionalAPIResponseModel is 工具服务商member组查询、新增接口 成功返回结果
type TaobaoTbkScMembergroupOptionalAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_membergroup_optional_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *TaobaoTbkScMembergroupOptionalMapData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScMembergroupOptionalAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScMembergroupOptionalAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScMembergroupOptionalAPIResponse)
	},
}

// GetTaobaoTbkScMembergroupOptionalAPIResponse 从 sync.Pool 获取 TaobaoTbkScMembergroupOptionalAPIResponse
func GetTaobaoTbkScMembergroupOptionalAPIResponse() *TaobaoTbkScMembergroupOptionalAPIResponse {
	return poolTaobaoTbkScMembergroupOptionalAPIResponse.Get().(*TaobaoTbkScMembergroupOptionalAPIResponse)
}

// ReleaseTaobaoTbkScMembergroupOptionalAPIResponse 将 TaobaoTbkScMembergroupOptionalAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScMembergroupOptionalAPIResponse(v *TaobaoTbkScMembergroupOptionalAPIResponse) {
	v.Reset()
	poolTaobaoTbkScMembergroupOptionalAPIResponse.Put(v)
}
