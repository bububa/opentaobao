package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxreportreportcrowdlistexpandAPIResponse 获取拓展人群数据报表 API返回值
// taobao.onebp.dkx.report.report.crowd.list.expand
//
// 获取拓展人群数据报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;effect&#34;:15,&#34;start_time&#34;:&#34;2021-09-08&#34;,&#34;end_time&#34;:&#34;2021-09-10&#34;,&#34;campaign_id_list&#34;:[2821811613],&#34;white_crowd_id_List&#34;:[12297883,12298696,12297989]}
type TaobaoonebpdkxreportreportcrowdlistexpandAPIResponse struct {
	model.CommonResponse
	TaobaoonebpdkxreportreportcrowdlistexpandAPIResponseModel
}

// TaobaoonebpdkxreportreportcrowdlistexpandAPIResponseModel is 获取拓展人群数据报表 成功返回结果
type TaobaoonebpdkxreportreportcrowdlistexpandAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_crowd_list_expand_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoonebpdkxreportreportcrowdlistexpandResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
