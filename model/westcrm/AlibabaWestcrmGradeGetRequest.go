package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取等级列表 APIRequest
alibaba.westcrm.grade.get

获取会员卡等级列表
*/
type AlibabaWestcrmGradeGetRequest struct {
    model.Params

    // 园区id
    campusId   int64 

}

func NewAlibabaWestcrmGradeGetRequest() *AlibabaWestcrmGradeGetRequest{
    return &AlibabaWestcrmGradeGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmGradeGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.grade.get"
}

func (r AlibabaWestcrmGradeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmGradeGetRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaWestcrmGradeGetRequest) GetCampusId() int64 {
    return r.campusId
}

