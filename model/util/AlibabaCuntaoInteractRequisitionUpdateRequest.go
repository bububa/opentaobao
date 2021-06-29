package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新物料制作状态 API请求
alibaba.cuntao.interact.requisition.update

村淘物料下沉，更新物料制作状态
*/
type AlibabaCuntaoInteractRequisitionUpdateRequest struct {
    model.Params
    // 物料制作状态
    status   string
    // 申请单id列表
    uuidList   []string
}

// 初始化AlibabaCuntaoInteractRequisitionUpdateRequest对象
func NewAlibabaCuntaoInteractRequisitionUpdateRequest() *AlibabaCuntaoInteractRequisitionUpdateRequest{
    return &AlibabaCuntaoInteractRequisitionUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCuntaoInteractRequisitionUpdateRequest) GetApiMethodName() string {
    return "alibaba.cuntao.interact.requisition.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCuntaoInteractRequisitionUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 物料制作状态
func (r *AlibabaCuntaoInteractRequisitionUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaCuntaoInteractRequisitionUpdateRequest) GetStatus() string {
    return r.status
}
// UuidList Setter
// 申请单id列表
func (r *AlibabaCuntaoInteractRequisitionUpdateRequest) SetUuidList(uuidList []string) error {
    r.uuidList = uuidList
    r.Set("uuid_list", uuidList)
    return nil
}

// UuidList Getter
func (r AlibabaCuntaoInteractRequisitionUpdateRequest) GetUuidList() []string {
    return r.uuidList
}
