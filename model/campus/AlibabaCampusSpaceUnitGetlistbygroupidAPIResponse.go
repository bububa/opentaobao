package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse
根据分组ID查询相应的空间单元 API返回值
alibaba.campus.space.unit.getlistbygroupid

根据分组ID查询相应的空间单元
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getListByGroupId */
type AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceUnitGetlistbygroupidAPIResponseModel
}

// AlibabaCampusSpaceUnitGetlistbygroupidAPIResponseModel is 根据分组ID查询相应的空间单元 成功返回结果
type AlibabaCampusSpaceUnitGetlistbygroupidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_unit_getlistbygroupid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}
