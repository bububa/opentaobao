package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse
空间分组id查业务属性实例 API返回值
alibaba.campus.space.group.getspacegroupwithattr

空间分组id查业务属性实例 */
type AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponseModel
}

// AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponseModel is 空间分组id查业务属性实例 成功返回结果
type AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_group_getspacegroupwithattr_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
