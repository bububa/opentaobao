package ihome

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实拍图发布审核状态查询API APIRequest
alibaba.ihome.ctom.content.publish.status

实拍图发布审核状态查询API
*/
type AlibabaIhomeCtomContentPublishStatusRequest struct {
    model.Params

    // 要查询投稿状态的ID列表
    idList   []int64 

}

func NewAlibabaIhomeCtomContentPublishStatusRequest() *AlibabaIhomeCtomContentPublishStatusRequest{
    return &AlibabaIhomeCtomContentPublishStatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIhomeCtomContentPublishStatusRequest) GetApiMethodName() string {
    return "alibaba.ihome.ctom.content.publish.status"
}

func (r AlibabaIhomeCtomContentPublishStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIhomeCtomContentPublishStatusRequest) SetIdList(idList []int64) error {
    r.idList = idList
    r.Set("id_list", idList)
    return nil
}

func (r AlibabaIhomeCtomContentPublishStatusRequest) GetIdList() []int64 {
    return r.idList
}

