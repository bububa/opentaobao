package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweikesubscinfogetAPIResponse 需求订单查询接口 API返回值
// taobao.weike.subscinfo.get
//
// 需求订单查询接口
type TaobaoweikesubscinfogetAPIResponse struct {
	model.CommonResponse
	TaobaoweikesubscinfogetAPIResponseModel
}

// TaobaoweikesubscinfogetAPIResponseModel is 需求订单查询接口 成功返回结果
type TaobaoweikesubscinfogetAPIResponseModel struct {
	XMLName xml.Name `xml:"weike_subscinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *SubscInfoWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
