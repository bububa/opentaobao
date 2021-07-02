package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetbyidAPIResponse 根据ID查询指定空间单元信息 API返回值
// alibaba.campus.space.unit.getbyid
//
// 根据ID查询指定空间单元信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getById
type AlibabaCampusSpaceUnitGetbyidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceUnitGetbyidAPIResponseModel
}

// AlibabaCampusSpaceUnitGetbyidAPIResponseModel is 根据ID查询指定空间单元信息 成功返回结果
type AlibabaCampusSpaceUnitGetbyidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_unit_getbyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
