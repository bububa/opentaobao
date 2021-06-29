package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动到店抽奖 API请求
alibaba.interact.isvlottery.idraw

互动到店抽奖
*/
type AlibabaInteractIsvlotteryIdrawRequest struct {
    model.Params
    // 互动实例ID
    _interactId   string
    // 抽奖ID列表 用逗号隔开
    _awardIds   string
    // 埋点参数
    _ua   string
}

// 初始化AlibabaInteractIsvlotteryIdrawRequest对象
func NewAlibabaInteractIsvlotteryIdrawRequest() *AlibabaInteractIsvlotteryIdrawRequest{
    return &AlibabaInteractIsvlotteryIdrawRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvlotteryIdrawRequest) GetApiMethodName() string {
    return "alibaba.interact.isvlottery.idraw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvlotteryIdrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InteractId Setter
// 互动实例ID
func (r *AlibabaInteractIsvlotteryIdrawRequest) SetInteractId(_interactId string) error {
    r._interactId = _interactId
    r.Set("interact_id", _interactId)
    return nil
}

// InteractId Getter
func (r AlibabaInteractIsvlotteryIdrawRequest) GetInteractId() string {
    return r._interactId
}
// AwardIds Setter
// 抽奖ID列表 用逗号隔开
func (r *AlibabaInteractIsvlotteryIdrawRequest) SetAwardIds(_awardIds string) error {
    r._awardIds = _awardIds
    r.Set("award_ids", _awardIds)
    return nil
}

// AwardIds Getter
func (r AlibabaInteractIsvlotteryIdrawRequest) GetAwardIds() string {
    return r._awardIds
}
// Ua Setter
// 埋点参数
func (r *AlibabaInteractIsvlotteryIdrawRequest) SetUa(_ua string) error {
    r._ua = _ua
    r.Set("ua", _ua)
    return nil
}

// Ua Getter
func (r AlibabaInteractIsvlotteryIdrawRequest) GetUa() string {
    return r._ua
}
