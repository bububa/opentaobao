package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取桌面升级任务 API请求
yunos.osupdate.appversion.query

分页获取桌面升级任务
*/
type YunosOsupdateAppversionQueryRequest struct {
    model.Params
    // 应用ID
    _appId   int64
    // 页码值
    _page   int64
    // 页大小
    _size   int64
}

// 初始化YunosOsupdateAppversionQueryRequest对象
func NewYunosOsupdateAppversionQueryRequest() *YunosOsupdateAppversionQueryRequest{
    return &YunosOsupdateAppversionQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionQueryRequest) GetApiMethodName() string {
    return "yunos.osupdate.appversion.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppId Setter
// 应用ID
func (r *YunosOsupdateAppversionQueryRequest) SetAppId(_appId int64) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r YunosOsupdateAppversionQueryRequest) GetAppId() int64 {
    return r._appId
}
// Page Setter
// 页码值
func (r *YunosOsupdateAppversionQueryRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r YunosOsupdateAppversionQueryRequest) GetPage() int64 {
    return r._page
}
// Size Setter
// 页大小
func (r *YunosOsupdateAppversionQueryRequest) SetSize(_size int64) error {
    r._size = _size
    r.Set("size", _size)
    return nil
}

// Size Getter
func (r YunosOsupdateAppversionQueryRequest) GetSize() int64 {
    return r._size
}
