package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpactivitieslistgetAPIResponse 营销活动列表查询 API返回值
// taobao.ump.activities.list.get
//
// 按照营销活动id的列表ids，查询对应的营销活动列表。
type TaobaoumpactivitieslistgetAPIResponse struct {
	model.CommonResponse
	TaobaoumpactivitieslistgetAPIResponseModel
}

// TaobaoumpactivitieslistgetAPIResponseModel is 营销活动列表查询 成功返回结果
type TaobaoumpactivitieslistgetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_activities_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销活动列表！
	Activities []string `json:"activities,omitempty" xml:"activities>string,omitempty"`
}
