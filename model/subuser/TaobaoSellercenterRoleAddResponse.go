package subuser

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
子账号角色的新增（指定卖家） APIResponse
taobao.sellercenter.role.add

给指定的卖家创建新的子账号角色<br/><br/>如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理<br/>例如：权限点列表如下<br/><br/>code=sell 宝贝管理<br/><br/>---------|code=sm 店铺管理<br/><br/>---------|---------|code=sm-design 如店铺装修<br/><br/>---------|---------|---------|code=sm-tbd-visit内店装修入口<br/><br/>---------|---------|---------|code=sm-tbd-publish内店装修发布<br/><br/>---------|---------|code=phone 手机淘宝店铺<br/><br/>调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色<br/><br/>code=sell 宝贝管理<br/><br/>---------|code=sm 店铺管理<br/><br/>---------|---------|code=sm-design 如店铺装修<br/><br/>---------|---------|---------|code=sm-tbd-visit内店装修入口<br/><br/>---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
*/
type TaobaoSellercenterRoleAddAPIResponse struct {
    model.CommonResponse
    TaobaoSellercenterRoleAddResponse
}

type TaobaoSellercenterRoleAddResponse struct {
    XMLName xml.Name `xml:"sellercenter_role_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 子账号角色
    
    Role   *Role `json:"role,omitempty" xml:"role,omitempty"`

    
}
