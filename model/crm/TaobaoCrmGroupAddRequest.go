package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家创建一个分组 API请求
taobao.crm.group.add

卖家创建一个新的分组，接口返回一个创建成功的分组的id
*/
type TaobaoCrmGroupAddRequest struct {
    model.Params
    // 分组名称，每个卖家最多可以拥有100个分组
    groupName   string
}

// 初始化TaobaoCrmGroupAddRequest对象
func NewTaobaoCrmGroupAddRequest() *TaobaoCrmGroupAddRequest{
    return &TaobaoCrmGroupAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupAddRequest) GetApiMethodName() string {
    return "taobao.crm.group.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupName Setter
// 分组名称，每个卖家最多可以拥有100个分组
func (r *TaobaoCrmGroupAddRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

// GroupName Getter
func (r TaobaoCrmGroupAddRequest) GetGroupName() string {
    return r.groupName
}
