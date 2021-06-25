package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新物料制作状态 APIRequest
alibaba.cuntao.interact.requisition.update

村淘物料下沉，更新物料制作状态
*/
type AlibabaCuntaoInteractRequisitionUpdateRequest struct {
    model.Params

    // 物料制作状态
    status   string 

    // 申请单id列表
    uuidList   []String 

}

func NewAlibabaCuntaoInteractRequisitionUpdateRequest() *AlibabaCuntaoInteractRequisitionUpdateRequest{
    return &AlibabaCuntaoInteractRequisitionUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCuntaoInteractRequisitionUpdateRequest) GetApiMethodName() string {
    return "alibaba.cuntao.interact.requisition.update"
}

func (r AlibabaCuntaoInteractRequisitionUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCuntaoInteractRequisitionUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaCuntaoInteractRequisitionUpdateRequest) GetStatus() string {
    return r.status
}

func (r *AlibabaCuntaoInteractRequisitionUpdateRequest) SetUuidList(uuidList []String) error {
    r.uuidList = uuidList
    r.Set("uuid_list", uuidList)
    return nil
}

func (r AlibabaCuntaoInteractRequisitionUpdateRequest) GetUuidList() []String {
    return r.uuidList
}

