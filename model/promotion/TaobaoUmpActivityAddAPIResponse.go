package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivityAddAPIResponse 新增优惠活动 API返回值
// taobao.ump.activity.add
//
// 新增优惠活动。设置优惠活动的基本信息，比如活动时间，活动针对的对象（可以是满足某些条件的买家）
type TaobaoUmpActivityAddAPIResponse struct {
	model.CommonResponse
	TaobaoUmpActivityAddAPIResponseModel
}

// TaobaoUmpActivityAddAPIResponseModel is 新增优惠活动 成功返回结果
type TaobaoUmpActivityAddAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_activity_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动id
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
}
