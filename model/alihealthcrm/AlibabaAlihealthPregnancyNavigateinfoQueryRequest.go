package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询底导数据 API请求
alibaba.alihealth.pregnancy.navigateinfo.query

备孕管理--获取底部导航信息
*/
type AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest struct {
    model.Params
    // 用户id
    _userId   int64
}

// 初始化AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest对象
func NewAlibabaAlihealthPregnancyNavigateinfoQueryRequest() *AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest{
    return &AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.navigateinfo.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户id
func (r *AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest) GetUserId() int64 {
    return r._userId
}
