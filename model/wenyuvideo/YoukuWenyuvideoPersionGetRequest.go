package wenyuvideo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据优酷人物ID获取人物详情页，包含相关影视和相关人物 API请求
youku.wenyuvideo.persion.get

根据优酷人物ID获取人物详情页，包含相关影视和相关人物
*/
type YoukuWenyuvideoPersionGetRequest struct {
    model.Params
    // 设备信息
    _systemInfo   string
    // 人物ID
    _personId   int64
}

// 初始化YoukuWenyuvideoPersionGetRequest对象
func NewYoukuWenyuvideoPersionGetRequest() *YoukuWenyuvideoPersionGetRequest{
    return &YoukuWenyuvideoPersionGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuWenyuvideoPersionGetRequest) GetApiMethodName() string {
    return "youku.wenyuvideo.persion.get"
}

// IRequest interface 方法, 获取API参数
func (r YoukuWenyuvideoPersionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SystemInfo Setter
// 设备信息
func (r *YoukuWenyuvideoPersionGetRequest) SetSystemInfo(_systemInfo string) error {
    r._systemInfo = _systemInfo
    r.Set("system_info", _systemInfo)
    return nil
}

// SystemInfo Getter
func (r YoukuWenyuvideoPersionGetRequest) GetSystemInfo() string {
    return r._systemInfo
}
// PersonId Setter
// 人物ID
func (r *YoukuWenyuvideoPersionGetRequest) SetPersonId(_personId int64) error {
    r._personId = _personId
    r.Set("person_id", _personId)
    return nil
}

// PersonId Getter
func (r YoukuWenyuvideoPersionGetRequest) GetPersonId() int64 {
    return r._personId
}
