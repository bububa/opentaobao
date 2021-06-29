package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商场列表 API请求
alibaba.westcrm.mall.list.get

根据园区id获取商场列表
*/
type AlibabaWestcrmMallListGetRequest struct {
    model.Params
    // 园区id
    campusId   int64
}

// 初始化AlibabaWestcrmMallListGetRequest对象
func NewAlibabaWestcrmMallListGetRequest() *AlibabaWestcrmMallListGetRequest{
    return &AlibabaWestcrmMallListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmMallListGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.mall.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmMallListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampusId Setter
// 园区id
func (r *AlibabaWestcrmMallListGetRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaWestcrmMallListGetRequest) GetCampusId() int64 {
    return r.campusId
}
