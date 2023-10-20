package crm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgroupappendAPIResponse 将一个分组添加到另外一个分组 API返回值
// taobao.crm.group.append
//
// 将某分组下的所有会员添加到另一个分组,注：1.该操作为异步任务，建议先调用taobao.crm.grouptask.check 确保涉及分组上没有任务；2.若分组下某会员分组数超最大限额，则该会员不会被添加到新分组，同时不影响其余会员添加分组，接口调用依然返回成功。
type TaobaocrmgroupappendAPIResponse struct {
	model.CommonResponse
	TaobaocrmgroupappendAPIResponseModel
}

// TaobaocrmgroupappendAPIResponseModel is 将一个分组添加到另外一个分组 成功返回结果
type TaobaocrmgroupappendAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_group_append_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步任务请求成功，添加任务是否完成通过taobao.crm.grouptask.check检测
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
