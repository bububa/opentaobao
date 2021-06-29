package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据关键词检索设备型号 API请求
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

// 初始化YunosOsupdateDeviceserviceSearchmodelsRequest对象
func NewYunosOsupdateDeviceserviceSearchmodelsRequest() *YunosOsupdateDeviceserviceSearchmodelsRequest{
    return &YunosOsupdateDeviceserviceSearchmodelsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateDeviceserviceSearchmodelsRequest) GetApiMethodName() string {
    return "yunos.osupdate.deviceservice.searchmodels"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateDeviceserviceSearchmodelsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParentId Setter
// 设备父ID
func (r *YunosOsupdateDeviceserviceSearchmodelsRequest) SetParentId(parentId int64) error {
    r.parentId = parentId
    r.Set("parent_id", parentId)
    return nil
}

// ParentId Getter
func (r YunosOsupdateDeviceserviceSearchmodelsRequest) GetParentId() int64 {
    return r.parentId
}
// Name Setter
// 关键词
func (r *YunosOsupdateDeviceserviceSearchmodelsRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r YunosOsupdateDeviceserviceSearchmodelsRequest) GetName() string {
    return r.name
}
