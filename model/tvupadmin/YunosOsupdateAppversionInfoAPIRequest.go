package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取应用升级详情 API请求
yunos.osupdate.appversion.info

获取应用升级详情
*/
type YunosOsupdateAppversionInfoAPIRequest struct {
    model.Params
    // 升级任务ID
    _id   int64
}

// 初始化YunosOsupdateAppversionInfoAPIRequest对象
func NewYunosOsupdateAppversionInfoRequest() *YunosOsupdateAppversionInfoAPIRequest{
    return &YunosOsupdateAppversionInfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionInfoAPIRequest) GetApiMethodName() string {
    return "yunos.osupdate.appversion.info"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionInfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 升级任务ID
func (r *YunosOsupdateAppversionInfoAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosOsupdateAppversionInfoAPIRequest) GetId() int64 {
    return r._id
}
