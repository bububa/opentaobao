package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台抽奖方案创建接口 API请求
alibaba.marketing.lottery.schema.create

抽奖平台抽奖方案创建接口
*/
type AlibabaMarketingLotterySchemaCreateRequest struct {
    model.Params
    // 创建抽奖方案请求对象
    _lotterySchemaCreate   *LotterySchemaCreateDto
}

// 初始化AlibabaMarketingLotterySchemaCreateRequest对象
func NewAlibabaMarketingLotterySchemaCreateRequest() *AlibabaMarketingLotterySchemaCreateRequest{
    return &AlibabaMarketingLotterySchemaCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotterySchemaCreateRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.schema.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotterySchemaCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotterySchemaCreate Setter
// 创建抽奖方案请求对象
func (r *AlibabaMarketingLotterySchemaCreateRequest) SetLotterySchemaCreate(_lotterySchemaCreate *LotterySchemaCreateDto) error {
    r._lotterySchemaCreate = _lotterySchemaCreate
    r.Set("lottery_schema_create", _lotterySchemaCreate)
    return nil
}

// LotterySchemaCreate Getter
func (r AlibabaMarketingLotterySchemaCreateRequest) GetLotterySchemaCreate() *LotterySchemaCreateDto {
    return r._lotterySchemaCreate
}
