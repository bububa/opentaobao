package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
空间单元id查业务属性实例 API返回值 
alibaba.campus.space.unit.getspaceunitwithattr

空间单元id查业务属性实例
*/
type AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponseModel
}

// 空间单元id查业务属性实例 成功返回结果
type AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_campus_space_unit_getspaceunitwithattr_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
