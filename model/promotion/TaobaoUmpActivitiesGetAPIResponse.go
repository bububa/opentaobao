package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpActivitiesGetAPIResponse
查询活动列表 API返回值
taobao.ump.activities.get

查询活动列表 */
type TaobaoUmpActivitiesGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpActivitiesGetAPIResponseModel
}

// TaobaoUmpActivitiesGetAPIResponseModel is 查询活动列表 成功返回结果
type TaobaoUmpActivitiesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_activities_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销活动内容，可以通过ump sdk来进行处理
	Contents []string `json:"contents,omitempty" xml:"contents>string,omitempty"`
	// 记录总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
