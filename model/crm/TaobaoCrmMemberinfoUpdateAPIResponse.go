package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑会员资料 API返回值 
taobao.crm.memberinfo.update

编辑会员的基本资料，接口返回会员信息修改是否成功
*/
type TaobaoCrmMemberinfoUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoCrmMemberinfoUpdateAPIResponseModel
}

// 编辑会员资料 成功返回结果
type TaobaoCrmMemberinfoUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"crm_memberinfo_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 会员信息修改是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
