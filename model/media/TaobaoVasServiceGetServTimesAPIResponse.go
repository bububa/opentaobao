package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaovasservicegetServTimesAPIResponse 查询某个用户图片空间的使用情况 API返回值
// taobao.vas.service.getServTimes
//
// 查询某个用户图片空间的使用情况
type TaobaovasservicegetServTimesAPIResponse struct {
	model.CommonResponse
	TaobaovasservicegetServTimesAPIResponseModel
}

// TaobaovasservicegetServTimesAPIResponseModel is 查询某个用户图片空间的使用情况 成功返回结果
type TaobaovasservicegetServTimesAPIResponseModel struct {
	XMLName xml.Name `xml:"vas_service_getServTimes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 总次数（容量）
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 剩余次数（容量）
	LeftNum int64 `json:"left_num,omitempty" xml:"left_num,omitempty"`
}
