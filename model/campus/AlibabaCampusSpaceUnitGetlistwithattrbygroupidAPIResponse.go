package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIResponse 根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】 API返回值
// alibaba.campus.space.unit.getlistwithattrbygroupid
//
// 根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】
type AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIResponseModel
}

// AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIResponseModel is 根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】 成功返回结果
type AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_unit_getlistwithattrbygroupid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
