package usergrowth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaousergrowthadmediadatasyncAPIResponse 媒体资源位投放效果数据回传 API返回值
// taobao.usergrowth.ad.media.data.sync
//
// 创意维度广告效果数据回传
type TaobaousergrowthadmediadatasyncAPIResponse struct {
	model.CommonResponse
	TaobaousergrowthadmediadatasyncAPIResponseModel
}

// TaobaousergrowthadmediadatasyncAPIResponseModel is 媒体资源位投放效果数据回传 成功返回结果
type TaobaousergrowthadmediadatasyncAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_ad_media_data_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 参数错误
	ResponseCode int64 `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 新增媒体数据成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 请求是否成功
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}
