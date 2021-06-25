package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
业务文件获取 APIRequest
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

func NewTaobaoFilesGetRequest() *TaobaoFilesGetRequest{
    return &TaobaoFilesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilesGetRequest) GetApiMethodName() string {
    return "taobao.files.get"
}

func (r TaobaoFilesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilesGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoFilesGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoFilesGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoFilesGetRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoFilesGetRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoFilesGetRequest) GetStatus() int64 {
    return r.status
}

