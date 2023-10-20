package usergrowth

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaogrowthreachingbuzzwordqueryAPIResponse 淘宝热词榜单数据查询接口 API返回值
// taobao.growth.reaching.buzzword.query
//
// 查询淘宝热词榜单数据
type TaobaogrowthreachingbuzzwordqueryAPIResponse struct {
	model.CommonResponse
	TaobaogrowthreachingbuzzwordqueryAPIResponseModel
}

// TaobaogrowthreachingbuzzwordqueryAPIResponseModel is 淘宝热词榜单数据查询接口 成功返回结果
type TaobaogrowthreachingbuzzwordqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"growth_reaching_buzzword_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 榜单词条列表
	Buzzwords []BuzzwordDto `json:"buzzwords,omitempty" xml:"buzzwords>buzzword_dto,omitempty"`
	// 榜单词条曝光链接
	ExposureUrl string `json:"exposure_url,omitempty" xml:"exposure_url,omitempty"`
	// 更多内容跳转链接
	Extra *ExtraDto `json:"extra,omitempty" xml:"extra,omitempty"`
	// 缓存时长，单位是秒
	CacheDuration int64 `json:"cache_duration,omitempty" xml:"cache_duration,omitempty"`
	// 是否针对此设备进行投放
	IsOffering bool `json:"is_offering,omitempty" xml:"is_offering,omitempty"`
}
