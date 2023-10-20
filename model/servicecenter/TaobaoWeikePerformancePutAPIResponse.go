package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweikeperformanceputAPIResponse 提交客服绩效接口 API返回值
// taobao.weike.performance.put
//
// 提交客服绩效接口
type TaobaoweikeperformanceputAPIResponse struct {
	model.CommonResponse
	TaobaoweikeperformanceputAPIResponseModel
}

// TaobaoweikeperformanceputAPIResponseModel is 提交客服绩效接口 成功返回结果
type TaobaoweikeperformanceputAPIResponseModel struct {
	XMLName xml.Name `xml:"weike_performance_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
