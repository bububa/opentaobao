package usergrowth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaousergrowthadmaterialdatasyncAPIResponse 素材投放效果数据回传 API返回值
// taobao.usergrowth.ad.material.data.sync
//
// 创意维度广告效果数据回传
type TaobaousergrowthadmaterialdatasyncAPIResponse struct {
	model.CommonResponse
	TaobaousergrowthadmaterialdatasyncAPIResponseModel
}

// TaobaousergrowthadmaterialdatasyncAPIResponseModel is 素材投放效果数据回传 成功返回结果
type TaobaousergrowthadmaterialdatasyncAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_ad_material_data_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	ResponseCode int64 `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 数据处理成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 请求是否成功
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}
