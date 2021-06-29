package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改一个已经存在的分组 API请求
taobao.crm.group.update

修改一个已经存在的分组，接口返回分组的修改是否成功
*/
type TaobaoCrmGroupUpdateRequest struct {
    model.Params
    // 分组的id
    groupId   int64
    // 新的分组名，分组名称不能包含|或者：
    newGroupName   string
}

// 初始化TaobaoCrmGroupUpdateRequest对象
func NewTaobaoCrmGroupUpdateRequest() *TaobaoCrmGroupUpdateRequest{
    return &TaobaoCrmGroupUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupUpdateRequest) GetApiMethodName() string {
    return "taobao.crm.group.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 分组的id
func (r *TaobaoCrmGroupUpdateRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r TaobaoCrmGroupUpdateRequest) GetGroupId() int64 {
    return r.groupId
}
// NewGroupName Setter
// 新的分组名，分组名称不能包含|或者：
func (r *TaobaoCrmGroupUpdateRequest) SetNewGroupName(newGroupName string) error {
    r.newGroupName = newGroupName
    r.Set("new_group_name", newGroupName)
    return nil
}

// NewGroupName Getter
func (r TaobaoCrmGroupUpdateRequest) GetNewGroupName() string {
    return r.newGroupName
}
