package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取桌面升级任务 APIRequest
yunos.osupdate.appversion.query

分页获取桌面升级任务
*/
type YunosOsupdateAppversionQueryRequest struct {
    model.Params

    // 应用ID
    appId   int64 

    // 页码值
    page   int64 

    // 页大小
    size   int64 

}

func NewYunosOsupdateAppversionQueryRequest() *YunosOsupdateAppversionQueryRequest{
    return &YunosOsupdateAppversionQueryRequest{
        Params: model.NewParams(),
    }
}

func (r YunosOsupdateAppversionQueryRequest) GetApiMethodName() string {
    return "yunos.osupdate.appversion.query"
}

func (r YunosOsupdateAppversionQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosOsupdateAppversionQueryRequest) SetAppId(appId int64) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r YunosOsupdateAppversionQueryRequest) GetAppId() int64 {
    return r.appId
}

func (r *YunosOsupdateAppversionQueryRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r YunosOsupdateAppversionQueryRequest) GetPage() int64 {
    return r.page
}

func (r *YunosOsupdateAppversionQueryRequest) SetSize(size int64) error {
    r.size = size
    r.Set("size", size)
    return nil
}

func (r YunosOsupdateAppversionQueryRequest) GetSize() int64 {
    return r.size
}

