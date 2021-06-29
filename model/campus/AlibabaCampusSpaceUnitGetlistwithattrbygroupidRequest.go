package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】 APIRequest
alibaba.campus.space.unit.getlistwithattrbygroupid

根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】
*/
type AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest struct {
    model.Params

    // 操作用户上下文
    context   *WorkBenchContext 

    // 分组id
    groupId   int64 

}

func NewAlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest() *AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest{
    return &AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) GetApiMethodName() string {
    return "alibaba.campus.space.unit.getlistwithattrbygroupid"
}

func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) GetContext() *WorkBenchContext {
    return r.context
}

func (r *AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest) GetGroupId() int64 {
    return r.groupId
}

