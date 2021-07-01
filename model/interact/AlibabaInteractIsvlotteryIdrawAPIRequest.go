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
type AlibabaInteractIsvlotteryIdrawAPIRequest struct {
    model.Params
    // 互动实例ID
    _interactId   string
    // 抽奖ID列表 用逗号隔开
    _awardIds   string
    // 埋点参数
    _ua   string
}

// 初始化AlibabaInteractIsvlotteryIdrawAPIRequest对象
func NewAlibabaInteractIsvlotteryIdrawRequest() *AlibabaInteractIsvlotteryIdrawAPIRequest{
    return &AlibabaInteractIsvlotteryIdrawAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvlotteryIdrawAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.isvlottery.idraw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvlotteryIdrawAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InteractId Setter
// 互动实例ID
func (r *AlibabaInteractIsvlotteryIdrawAPIRequest) SetInteractId(_interactId string) error {
    r._interactId = _interactId
    r.Set("interact_id", _interactId)
    return nil
}

// InteractId Getter
func (r AlibabaInteractIsvlotteryIdrawAPIRequest) GetInteractId() string {
    return r._interactId
}
// AwardIds Setter
// 抽奖ID列表 用逗号隔开
func (r *AlibabaInteractIsvlotteryIdrawAPIRequest) SetAwardIds(_awardIds string) error {
    r._awardIds = _awardIds
    r.Set("award_ids", _awardIds)
    return nil
}

// AwardIds Getter
func (r AlibabaInteractIsvlotteryIdrawAPIRequest) GetAwardIds() string {
    return r._awardIds
}
// Ua Setter
// 埋点参数
func (r *AlibabaInteractIsvlotteryIdrawAPIRequest) SetUa(_ua string) error {
    r._ua = _ua
    r.Set("ua", _ua)
    return nil
}

// Ua Getter
func (r AlibabaInteractIsvlotteryIdrawAPIRequest) GetUa() string {
    return r._ua
}
