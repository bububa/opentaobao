package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotterySchemaCreateAPIRequest 抽奖平台抽奖方案创建接口 API请求
// alibaba.marketing.lottery.schema.create
//
// 抽奖平台抽奖方案创建接口
type AlibabaMarketingLotterySchemaCreateAPIRequest struct {
	model.Params
	// 创建抽奖方案请求对象
	_lotterySchemaCreate *LotterySchemaCreateDto
}

// NewAlibabaMarketingLotterySchemaCreateRequest 初始化AlibabaMarketingLotterySchemaCreateAPIRequest对象
func NewAlibabaMarketingLotterySchemaCreateRequest() *AlibabaMarketingLotterySchemaCreateAPIRequest {
	return &AlibabaMarketingLotterySchemaCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotterySchemaCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.marketing.lottery.schema.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotterySchemaCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LotterySchemaCreate Setter
// 创建抽奖方案请求对象
func (r *AlibabaMarketingLotterySchemaCreateAPIRequest) SetLotterySchemaCreate(_lotterySchemaCreate *LotterySchemaCreateDto) error {
	r._lotterySchemaCreate = _lotterySchemaCreate
	r.Set("lottery_schema_create", _lotterySchemaCreate)
	return nil
}

// Get LotterySchemaCreate Getter
func (r AlibabaMarketingLotterySchemaCreateAPIRequest) GetLotterySchemaCreate() *LotterySchemaCreateDto {
	return r._lotterySchemaCreate
}
