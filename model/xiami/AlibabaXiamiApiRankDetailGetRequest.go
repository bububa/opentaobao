package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
排行榜详情 APIRequest
alibaba.xiami.api.rank.detail.get

虾米排行榜详情数据
*/
type AlibabaXiamiApiRankDetailGetRequest struct {
    model.Params

    // 榜单ID
    billboardId   int64 

    // 调用来源
    bizCode   string 

}

func NewAlibabaXiamiApiRankDetailGetRequest() *AlibabaXiamiApiRankDetailGetRequest{
    return &AlibabaXiamiApiRankDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiRankDetailGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.rank.detail.get"
}

func (r AlibabaXiamiApiRankDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiRankDetailGetRequest) SetBillboardId(billboardId int64) error {
    r.billboardId = billboardId
    r.Set("billboard_id", billboardId)
    return nil
}

func (r AlibabaXiamiApiRankDetailGetRequest) GetBillboardId() int64 {
    return r.billboardId
}

func (r *AlibabaXiamiApiRankDetailGetRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r AlibabaXiamiApiRankDetailGetRequest) GetBizCode() string {
    return r.bizCode
}

