package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionMealGetAPIResponse
搭配套餐查询 API返回值
taobao.promotion.meal.get

搭配套餐查询。每个卖家最多创建50个搭配套餐，所以查询不会分页，会将所有的满足状态的搭配套餐全部查出。该接口不会校验商品的下架或库存为0，查询结果的状态表明搭配套餐在数据库中的状态，商品的状态请isv自己验证。在卖家后台页面点击查看会触发数据库状态的修改。 */
type TaobaoPromotionMealGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionMealGetAPIResponseModel
}

// TaobaoPromotionMealGetAPIResponseModel is 搭配套餐查询 成功返回结果
type TaobaoPromotionMealGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_meal_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搭配套餐列表。
	MealList []Meal `json:"meal_list,omitempty" xml:"meal_list>meal,omitempty"`
}
