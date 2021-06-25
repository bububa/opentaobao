package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动到店抽奖 APIRequest
alibaba.interact.isvlottery.idraw

互动到店抽奖
*/
type AlibabaInteractIsvlotteryIdrawRequest struct {
    model.Params

    // 互动实例ID
    interactId   string 

    // 抽奖ID列表 用逗号隔开
    awardIds   string 

    // 埋点参数
    ua   string 

}

func NewAlibabaInteractIsvlotteryIdrawRequest() *AlibabaInteractIsvlotteryIdrawRequest{
    return &AlibabaInteractIsvlotteryIdrawRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractIsvlotteryIdrawRequest) GetApiMethodName() string {
    return "alibaba.interact.isvlottery.idraw"
}

func (r AlibabaInteractIsvlotteryIdrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractIsvlotteryIdrawRequest) SetInteractId(interactId string) error {
    r.interactId = interactId
    r.Set("interact_id", interactId)
    return nil
}

func (r AlibabaInteractIsvlotteryIdrawRequest) GetInteractId() string {
    return r.interactId
}

func (r *AlibabaInteractIsvlotteryIdrawRequest) SetAwardIds(awardIds string) error {
    r.awardIds = awardIds
    r.Set("award_ids", awardIds)
    return nil
}

func (r AlibabaInteractIsvlotteryIdrawRequest) GetAwardIds() string {
    return r.awardIds
}

func (r *AlibabaInteractIsvlotteryIdrawRequest) SetUa(ua string) error {
    r.ua = ua
    r.Set("ua", ua)
    return nil
}

func (r AlibabaInteractIsvlotteryIdrawRequest) GetUa() string {
    return r.ua
}

