package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallcarxcarsynchronizecarlinedataAPIResponse 我的爱卡车型配置数据 API返回值
// tmall.car.xcar.synchronize.car.line.data
//
// 同步我的爱卡车系配置数据到天猫汽车
type TmallcarxcarsynchronizecarlinedataAPIResponse struct {
	model.CommonResponse
	TmallcarxcarsynchronizecarlinedataAPIResponseModel
}

// TmallcarxcarsynchronizecarlinedataAPIResponseModel is 我的爱卡车型配置数据 成功返回结果
type TmallcarxcarsynchronizecarlinedataAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_xcar_synchronize_car_line_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象总体信息
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
