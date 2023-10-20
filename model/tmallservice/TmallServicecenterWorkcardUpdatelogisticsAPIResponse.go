package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardupdatelogisticsAPIResponse 更新物流进度 API返回值
// tmall.servicecenter.workcard.updatelogistics
//
// 提供给外部合作服务商的物流进度更改接口
type TmallservicecenterworkcardupdatelogisticsAPIResponse struct {
	model.CommonResponse
	TmallservicecenterworkcardupdatelogisticsAPIResponseModel
}

// TmallservicecenterworkcardupdatelogisticsAPIResponseModel is 更新物流进度 成功返回结果
type TmallservicecenterworkcardupdatelogisticsAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_updatelogistics_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
