package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据关键词检索设备型号 APIRequest
yunos.osupdate.deviceservice.searchmodels

根据关键词检索设备型号
*/
type YunosOsupdateDeviceserviceSearchmodelsRequest struct {
    model.Params

    // 设备父ID
    parentId   int64 

    // 关键词
    name   string 

}

func NewYunosOsupdateDeviceserviceSearchmodelsRequest() *YunosOsupdateDeviceserviceSearchmodelsRequest{
    return &YunosOsupdateDeviceserviceSearchmodelsRequest{
        Params: model.NewParams(),
    }
}

func (r YunosOsupdateDeviceserviceSearchmodelsRequest) GetApiMethodName() string {
    return "yunos.osupdate.deviceservice.searchmodels"
}

func (r YunosOsupdateDeviceserviceSearchmodelsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosOsupdateDeviceserviceSearchmodelsRequest) SetParentId(parentId int64) error {
    r.parentId = parentId
    r.Set("parent_id", parentId)
    return nil
}

func (r YunosOsupdateDeviceserviceSearchmodelsRequest) GetParentId() int64 {
    return r.parentId
}

func (r *YunosOsupdateDeviceserviceSearchmodelsRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r YunosOsupdateDeviceserviceSearchmodelsRequest) GetName() string {
    return r.name
}

