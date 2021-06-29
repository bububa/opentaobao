package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询底导数据 APIRequest
alibaba.alihealth.pregnancy.navigateinfo.query

备孕管理--获取底部导航信息
*/
type AlibabaAlihealthPregnancyNavigateinfoQueryRequest struct {
    model.Params

    // 用户id
    userId   int64 

}

func NewAlibabaAlihealthPregnancyNavigateinfoQueryRequest() *AlibabaAlihealthPregnancyNavigateinfoQueryRequest{
    return &AlibabaAlihealthPregnancyNavigateinfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthPregnancyNavigateinfoQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.navigateinfo.query"
}

func (r AlibabaAlihealthPregnancyNavigateinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthPregnancyNavigateinfoQueryRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlihealthPregnancyNavigateinfoQueryRequest) GetUserId() int64 {
    return r.userId
}

