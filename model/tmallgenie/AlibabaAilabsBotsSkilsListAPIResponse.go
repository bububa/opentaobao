package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsBotsSkilsListAPIResponse 对外设备获取技能列表 API返回值
// alibaba.ailabs.bots.skils.list
//
// 获取ai开放平台技能列表
type AlibabaAilabsBotsSkilsListAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsBotsSkilsListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsBotsSkilsListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsBotsSkilsListAPIResponseModel).Reset()
}

// AlibabaAilabsBotsSkilsListAPIResponseModel is 对外设备获取技能列表 成功返回结果
type AlibabaAilabsBotsSkilsListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_bots_skils_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// {        &#34;skillId&#34;: 209,        &#34;invocationName&#34;: &#34;中文先生&#34;,        &#34;name&#34;: &#34;测试34&#34;,        &#34;serviceProviders&#34;: [          {            &#34;icon&#34;: &#34;//arplatform.alicdn.com/images/3/1498910818259.png&#34;,            &#34;name&#34;: &#34;provider1&#34;,          }        ],        &#34;botId&#34;: 10,        &#34;iconImgUrl&#34;: &#34;//arplatform.alicdn.com/images/244/1501764397807.png&#34;,        &#34;longDesc&#34;: &#34;中文先生是学中文的好帮手。查中文、查成语、听故事样样行。&#34;      }
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsBotsSkilsListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsBotsSkilsListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsBotsSkilsListAPIResponse)
	},
}

// GetAlibabaAilabsBotsSkilsListAPIResponse 从 sync.Pool 获取 AlibabaAilabsBotsSkilsListAPIResponse
func GetAlibabaAilabsBotsSkilsListAPIResponse() *AlibabaAilabsBotsSkilsListAPIResponse {
	return poolAlibabaAilabsBotsSkilsListAPIResponse.Get().(*AlibabaAilabsBotsSkilsListAPIResponse)
}

// ReleaseAlibabaAilabsBotsSkilsListAPIResponse 将 AlibabaAilabsBotsSkilsListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsBotsSkilsListAPIResponse(v *AlibabaAilabsBotsSkilsListAPIResponse) {
	v.Reset()
	poolAlibabaAilabsBotsSkilsListAPIResponse.Put(v)
}
