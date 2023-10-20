package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpAdzoneFindconfiglistAPIResponse 查询所有可用资源包信息 API返回值
// taobao.universalbp.adzone.findconfiglist
//
// 查询该场景下，所有可用的资源包，可能存在数据重复。但是针对不同子场景和推广设置，可以选用的资源包有差异，建议关注补充文档，或者根据bp前端的限制，进行传参。
type TaobaoUniversalbpAdzoneFindconfiglistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpAdzoneFindconfiglistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpAdzoneFindconfiglistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpAdzoneFindconfiglistAPIResponseModel).Reset()
}

// TaobaoUniversalbpAdzoneFindconfiglistAPIResponseModel is 查询所有可用资源包信息 成功返回结果
type TaobaoUniversalbpAdzoneFindconfiglistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_adzone_findconfiglist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpAdzoneFindconfiglistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpAdzoneFindconfiglistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpAdzoneFindconfiglistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpAdzoneFindconfiglistAPIResponse)
	},
}

// GetTaobaoUniversalbpAdzoneFindconfiglistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpAdzoneFindconfiglistAPIResponse
func GetTaobaoUniversalbpAdzoneFindconfiglistAPIResponse() *TaobaoUniversalbpAdzoneFindconfiglistAPIResponse {
	return poolTaobaoUniversalbpAdzoneFindconfiglistAPIResponse.Get().(*TaobaoUniversalbpAdzoneFindconfiglistAPIResponse)
}

// ReleaseTaobaoUniversalbpAdzoneFindconfiglistAPIResponse 将 TaobaoUniversalbpAdzoneFindconfiglistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpAdzoneFindconfiglistAPIResponse(v *TaobaoUniversalbpAdzoneFindconfiglistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpAdzoneFindconfiglistAPIResponse.Put(v)
}
