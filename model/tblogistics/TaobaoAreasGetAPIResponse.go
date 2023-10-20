package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAreasGetAPIResponse 查询地址区域 API返回值
// taobao.areas.get
//
// 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
// &lt;a href=&#34;http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/&#34; target=&#34;_blank&#34;&gt; http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/&lt;/a&gt;
type TaobaoAreasGetAPIResponse struct {
	model.CommonResponse
	TaobaoAreasGetAPIResponseModel
}

// TaobaoAreasGetAPIResponseModel is 查询地址区域 成功返回结果
type TaobaoAreasGetAPIResponseModel struct {
	XMLName xml.Name `xml:"areas_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 地址区域信息列表.返回的Area包含的具体信息为入参fields请求的字段信息 。
	Areas []Area `json:"areas,omitempty" xml:"areas>area,omitempty"`
}
