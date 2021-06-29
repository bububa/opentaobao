package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
业务文件获取 API请求
taobao.files.get

获取业务方暂存给ISV的文件列表
*/
type TaobaoFilesGetRequest struct {
    model.Params
    // 搜索开始时间
    startDate   string
    // 搜索结束时间
    endDate   string
    // 下载链接状态。1:未下载。2:已下载
    status   int64
}

// 初始化TaobaoFilesGetRequest对象
func NewTaobaoFilesGetRequest() *TaobaoFilesGetRequest{
    return &TaobaoFilesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFilesGetRequest) GetApiMethodName() string {
    return "taobao.files.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFilesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 搜索开始时间
func (r *TaobaoFilesGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoFilesGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 搜索结束时间
func (r *TaobaoFilesGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoFilesGetRequest) GetEndDate() string {
    return r.endDate
}
// Status Setter
// 下载链接状态。1:未下载。2:已下载
func (r *TaobaoFilesGetRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoFilesGetRequest) GetStatus() int64 {
    return r.status
}
