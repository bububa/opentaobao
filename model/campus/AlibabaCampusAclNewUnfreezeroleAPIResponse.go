package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewUnfreezeroleAPIResponse
解冻角色 API返回值
alibaba.campus.acl.new.unfreezerole

解冻角色 */
type AlibabaCampusAclNewUnfreezeroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewUnfreezeroleAPIResponseModel
}

// AlibabaCampusAclNewUnfreezeroleAPIResponseModel is 解冻角色 成功返回结果
type AlibabaCampusAclNewUnfreezeroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_unfreezerole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
