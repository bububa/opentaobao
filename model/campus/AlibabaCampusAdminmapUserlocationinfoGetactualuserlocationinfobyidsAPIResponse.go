package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusadminmapuserlocationinfogetactualuserlocationinfobyidsAPIResponse 根据userId(支持单个或批量)获取用户实时位置信息 API返回值
// alibaba.campus.adminmap.userlocationinfo.getactualuserlocationinfobyids
//
// 根据userId(支持单个或批量)获取用户实时位置信息
// HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
// HSF方法名称：getActualUserLocationInfoByIds
type AlibabacampusadminmapuserlocationinfogetactualuserlocationinfobyidsAPIResponse struct {
	model.CommonResponse
	AlibabacampusadminmapuserlocationinfogetactualuserlocationinfobyidsAPIResponseModel
}

// AlibabacampusadminmapuserlocationinfogetactualuserlocationinfobyidsAPIResponseModel is 根据userId(支持单个或批量)获取用户实时位置信息 成功返回结果
type AlibabacampusadminmapuserlocationinfogetactualuserlocationinfobyidsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_adminmap_userlocationinfo_getactualuserlocationinfobyids_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
