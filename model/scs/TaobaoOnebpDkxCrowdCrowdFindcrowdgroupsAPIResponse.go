package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIResponse 获取人群组 API返回值
// taobao.onebp.dkx.crowd.crowd.findcrowdgroups
//
// 获取人群组
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{ &#34;market_scene&#34;: &#34;ad_strategy_laxin&#34;}
type TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIResponseModel
}

// TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIResponseModel is 获取人群组 成功返回结果
type TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_crowd_crowd_findcrowdgroups_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
