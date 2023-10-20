package smartstore

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoIstoreAreasGetAPIResponse 智慧门店区域编码查询 API返回值
// taobao.istore.areas.get
//
// 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
// &lt;a href=&#34;http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html&#34;&gt;http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html&lt;/a&gt;
type TaobaoIstoreAreasGetAPIResponse struct {
	model.CommonResponse
	TaobaoIstoreAreasGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoIstoreAreasGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoIstoreAreasGetAPIResponseModel).Reset()
}

// TaobaoIstoreAreasGetAPIResponseModel is 智慧门店区域编码查询 成功返回结果
type TaobaoIstoreAreasGetAPIResponseModel struct {
	XMLName xml.Name `xml:"istore_areas_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 地址区域信息列表.返回的Area包含的具体信息为入参fields请求的字段信息 。
	Areas []Area `json:"areas,omitempty" xml:"areas>area,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoIstoreAreasGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Areas = m.Areas[:0]
}

var poolTaobaoIstoreAreasGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoIstoreAreasGetAPIResponse)
	},
}

// GetTaobaoIstoreAreasGetAPIResponse 从 sync.Pool 获取 TaobaoIstoreAreasGetAPIResponse
func GetTaobaoIstoreAreasGetAPIResponse() *TaobaoIstoreAreasGetAPIResponse {
	return poolTaobaoIstoreAreasGetAPIResponse.Get().(*TaobaoIstoreAreasGetAPIResponse)
}

// ReleaseTaobaoIstoreAreasGetAPIResponse 将 TaobaoIstoreAreasGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoIstoreAreasGetAPIResponse(v *TaobaoIstoreAreasGetAPIResponse) {
	v.Reset()
	poolTaobaoIstoreAreasGetAPIResponse.Put(v)
}
