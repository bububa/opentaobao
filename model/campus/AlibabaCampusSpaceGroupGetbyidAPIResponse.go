package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceGroupGetbyidAPIResponse
根据分组ID查询相关的空间分组信息 API返回值
alibaba.campus.space.group.getbyid

根据分组ID查询相关的空间分组信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getById */
type AlibabaCampusSpaceGroupGetbyidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceGroupGetbyidAPIResponseModel
}

// AlibabaCampusSpaceGroupGetbyidAPIResponseModel is 根据分组ID查询相关的空间分组信息 成功返回结果
type AlibabaCampusSpaceGroupGetbyidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_group_getbyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
