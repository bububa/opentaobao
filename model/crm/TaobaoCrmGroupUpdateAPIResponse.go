package crm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgroupupdateAPIResponse 修改一个已经存在的分组 API返回值
// taobao.crm.group.update
//
// 修改一个已经存在的分组，接口返回分组的修改是否成功
type TaobaocrmgroupupdateAPIResponse struct {
	model.CommonResponse
	TaobaocrmgroupupdateAPIResponseModel
}

// TaobaocrmgroupupdateAPIResponseModel is 修改一个已经存在的分组 成功返回结果
type TaobaocrmgroupupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_group_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分组修改是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
