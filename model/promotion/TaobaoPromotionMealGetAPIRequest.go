package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionMealGetAPIRequest 搭配套餐查询 API请求
// taobao.promotion.meal.get
//
// 搭配套餐查询。每个卖家最多创建50个搭配套餐，所以查询不会分页，会将所有的满足状态的搭配套餐全部查出。该接口不会校验商品的下架或库存为0，查询结果的状态表明搭配套餐在数据库中的状态，商品的状态请isv自己验证。在卖家后台页面点击查看会触发数据库状态的修改。
type TaobaoPromotionMealGetAPIRequest struct {
	model.Params
	// 套餐状态。有效：VALID;失效：INVALID(有效套餐为可使用的套餐,无效套餐为套餐中有商品下架或库存为0时)。默认时两种情况都会查询。
	_status string
	// 搭配套餐id
	_mealId int64
}

// NewTaobaoPromotionMealGetRequest 初始化TaobaoPromotionMealGetAPIRequest对象
func NewTaobaoPromotionMealGetRequest() *TaobaoPromotionMealGetAPIRequest {
	return &TaobaoPromotionMealGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionMealGetAPIRequest) Reset() {
	r._status = ""
	r._mealId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionMealGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.meal.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionMealGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionMealGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 套餐状态。有效：VALID;失效：INVALID(有效套餐为可使用的套餐,无效套餐为套餐中有商品下架或库存为0时)。默认时两种情况都会查询。
func (r *TaobaoPromotionMealGetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoPromotionMealGetAPIRequest) GetStatus() string {
	return r._status
}

// SetMealId is MealId Setter
// 搭配套餐id
func (r *TaobaoPromotionMealGetAPIRequest) SetMealId(_mealId int64) error {
	r._mealId = _mealId
	r.Set("meal_id", _mealId)
	return nil
}

// GetMealId MealId Getter
func (r TaobaoPromotionMealGetAPIRequest) GetMealId() int64 {
	return r._mealId
}

var poolTaobaoPromotionMealGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionMealGetRequest()
	},
}

// GetTaobaoPromotionMealGetRequest 从 sync.Pool 获取 TaobaoPromotionMealGetAPIRequest
func GetTaobaoPromotionMealGetAPIRequest() *TaobaoPromotionMealGetAPIRequest {
	return poolTaobaoPromotionMealGetAPIRequest.Get().(*TaobaoPromotionMealGetAPIRequest)
}

// ReleaseTaobaoPromotionMealGetAPIRequest 将 TaobaoPromotionMealGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionMealGetAPIRequest(v *TaobaoPromotionMealGetAPIRequest) {
	v.Reset()
	poolTaobaoPromotionMealGetAPIRequest.Put(v)
}
