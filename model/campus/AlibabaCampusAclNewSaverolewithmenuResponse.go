package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
保存角色级联保存角色和权限的关系 APIResponse
alibaba.campus.acl.new.saverolewithmenu

保存角色级联保存角色和权限的关系
*/
type AlibabaCampusAclNewSaverolewithmenuAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewSaverolewithmenuResponse
}

type AlibabaCampusAclNewSaverolewithmenuResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_saverolewithmenu_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回结果
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
