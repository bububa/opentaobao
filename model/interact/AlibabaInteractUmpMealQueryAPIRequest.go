package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractUmpMealQueryAPIRequest 淘宝卖家搭配套餐查询 API请求
// alibaba.interact.ump.meal.query
//
// 查询卖家在优惠平台设置的搭配套餐列表，每个套餐包括名称、套餐价格、手淘套餐购买链接
type AlibabaInteractUmpMealQueryAPIRequest struct {
	model.Params
}

// NewAlibabaInteractUmpMealQueryRequest 初始化AlibabaInteractUmpMealQueryAPIRequest对象
func NewAlibabaInteractUmpMealQueryRequest() *AlibabaInteractUmpMealQueryAPIRequest {
	return &AlibabaInteractUmpMealQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractUmpMealQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractUmpMealQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.ump.meal.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractUmpMealQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractUmpMealQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractUmpMealQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractUmpMealQueryRequest()
	},
}

// GetAlibabaInteractUmpMealQueryRequest 从 sync.Pool 获取 AlibabaInteractUmpMealQueryAPIRequest
func GetAlibabaInteractUmpMealQueryAPIRequest() *AlibabaInteractUmpMealQueryAPIRequest {
	return poolAlibabaInteractUmpMealQueryAPIRequest.Get().(*AlibabaInteractUmpMealQueryAPIRequest)
}

// ReleaseAlibabaInteractUmpMealQueryAPIRequest 将 AlibabaInteractUmpMealQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractUmpMealQueryAPIRequest(v *AlibabaInteractUmpMealQueryAPIRequest) {
	v.Reset()
	poolAlibabaInteractUmpMealQueryAPIRequest.Put(v)
}
