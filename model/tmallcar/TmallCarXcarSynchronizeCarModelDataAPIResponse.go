package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallcarxcarsynchronizecarmodeldataAPIResponse 爱车车型数据同步 API返回值
// tmall.car.xcar.synchronize.car.model.data
//
// 爱车汽车车型数据同步到天猫
type TmallcarxcarsynchronizecarmodeldataAPIResponse struct {
	model.CommonResponse
	TmallcarxcarsynchronizecarmodeldataAPIResponseModel
}

// TmallcarxcarsynchronizecarmodeldataAPIResponseModel is 爱车车型数据同步 成功返回结果
type TmallcarxcarsynchronizecarmodeldataAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_xcar_synchronize_car_model_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象描述
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
