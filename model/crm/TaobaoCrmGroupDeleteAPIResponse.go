package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除分组 API返回值 
taobao.crm.group.delete

将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
*/
type TaobaoCrmGroupDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoCrmGroupDeleteAPIResponseModel
}

// 删除分组 成功返回结果
type TaobaoCrmGroupDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"crm_group_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步任务请求成功，是否执行完毕需要通过taobao.crm.grouptask.check检测
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
