package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取等级列表 API请求
alibaba.westcrm.grade.get

获取会员卡等级列表
*/
type AlibabaWestcrmGradeGetRequest struct {
    model.Params
    // 园区id
    campusId   int64
}

// 初始化AlibabaWestcrmGradeGetRequest对象
func NewAlibabaWestcrmGradeGetRequest() *AlibabaWestcrmGradeGetRequest{
    return &AlibabaWestcrmGradeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmGradeGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.grade.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmGradeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampusId Setter
// 园区id
func (r *AlibabaWestcrmGradeGetRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaWestcrmGradeGetRequest) GetCampusId() int64 {
    return r.campusId
}
