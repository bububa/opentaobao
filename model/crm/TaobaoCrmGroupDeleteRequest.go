package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除分组 API请求
taobao.crm.group.delete

将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
*/
type TaobaoCrmGroupDeleteRequest struct {
    model.Params
    // 要删除的分组id
    _groupId   int64
}

// 初始化TaobaoCrmGroupDeleteRequest对象
func NewTaobaoCrmGroupDeleteRequest() *TaobaoCrmGroupDeleteRequest{
    return &TaobaoCrmGroupDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupDeleteRequest) GetApiMethodName() string {
    return "taobao.crm.group.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 要删除的分组id
func (r *TaobaoCrmGroupDeleteRequest) SetGroupId(_groupId int64) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r TaobaoCrmGroupDeleteRequest) GetGroupId() int64 {
    return r._groupId
}
