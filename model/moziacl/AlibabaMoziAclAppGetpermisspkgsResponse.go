package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取应用的权限套餐 APIResponse
alibaba.mozi.acl.app.getpermisspkgs

分页查询应用下的权限套餐列表
*/
type AlibabaMoziAclAppGetpermisspkgsAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclAppGetpermisspkgsResponse
}

type AlibabaMoziAclAppGetpermisspkgsResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_app_getpermisspkgs_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 获取应用的权限套餐列表结果对象
    
    Result   *AppPermissionPackageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
