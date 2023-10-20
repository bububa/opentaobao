package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIResponse 上传用户实时位置 API返回值
// alibaba.campus.adminmap.userlocationinfo.insertactualuserlocationinfo
//
// 上传用户实时位置
// HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
// HSF方法名称：insertActualUserLocationInfo
type AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIResponse struct {
	model.CommonResponse
	AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIResponseModel
}

// AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIResponseModel is 上传用户实时位置 成功返回结果
type AlibabacampusadminmapuserlocationinfoinsertactualuserlocationinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_adminmap_userlocationinfo_insertactualuserlocationinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
