package usergrowth

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaogrowthreachingbrowsersearchAPIResponse 查询搜索关联 API返回值
// taobao.growth.reaching.browser.search
//
// 查询搜索关联
type TaobaogrowthreachingbrowsersearchAPIResponse struct {
	model.CommonResponse
	TaobaogrowthreachingbrowsersearchAPIResponseModel
}

// TaobaogrowthreachingbrowsersearchAPIResponseModel is 查询搜索关联 成功返回结果
type TaobaogrowthreachingbrowsersearchAPIResponseModel struct {
	XMLName xml.Name `xml:"growth_reaching_browser_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// []
	Products []ProductDto `json:"products,omitempty" xml:"products>product_dto,omitempty"`
	// 榜单曝光链接
	ExposureUrl string `json:"exposure_url,omitempty" xml:"exposure_url,omitempty"`
	// 是否参竞
	Offering bool `json:"offering,omitempty" xml:"offering,omitempty"`
}
