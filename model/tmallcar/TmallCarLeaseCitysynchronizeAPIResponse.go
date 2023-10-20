package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleasecitysynchronizeAPIResponse 天猫开新车租后分期城市信息同步 API返回值
// tmall.car.lease.citysynchronize
//
// 天猫开新车租后分期城市信息同步
type TmallcarleasecitysynchronizeAPIResponse struct {
	model.CommonResponse
	TmallcarleasecitysynchronizeAPIResponseModel
}

// TmallcarleasecitysynchronizeAPIResponseModel is 天猫开新车租后分期城市信息同步 成功返回结果
type TmallcarleasecitysynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_citysynchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
