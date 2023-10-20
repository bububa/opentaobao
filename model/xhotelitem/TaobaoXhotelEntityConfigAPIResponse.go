package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelEntityConfigAPIResponse 飞猪商品各实体通用配置 API返回值
// taobao.xhotel.entity.config
//
// 飞猪商品各实体通用配置服务
type TaobaoXhotelEntityConfigAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelEntityConfigAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelEntityConfigAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelEntityConfigAPIResponseModel).Reset()
}

// TaobaoXhotelEntityConfigAPIResponseModel is 飞猪商品各实体通用配置 成功返回结果
type TaobaoXhotelEntityConfigAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_entity_config_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelEntityConfigAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolTaobaoXhotelEntityConfigAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelEntityConfigAPIResponse)
	},
}

// GetTaobaoXhotelEntityConfigAPIResponse 从 sync.Pool 获取 TaobaoXhotelEntityConfigAPIResponse
func GetTaobaoXhotelEntityConfigAPIResponse() *TaobaoXhotelEntityConfigAPIResponse {
	return poolTaobaoXhotelEntityConfigAPIResponse.Get().(*TaobaoXhotelEntityConfigAPIResponse)
}

// ReleaseTaobaoXhotelEntityConfigAPIResponse 将 TaobaoXhotelEntityConfigAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelEntityConfigAPIResponse(v *TaobaoXhotelEntityConfigAPIResponse) {
	v.Reset()
	poolTaobaoXhotelEntityConfigAPIResponse.Put(v)
}
