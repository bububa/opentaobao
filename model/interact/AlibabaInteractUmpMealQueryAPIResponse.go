package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractUmpMealQueryAPIResponse 淘宝卖家搭配套餐查询 API返回值
// alibaba.interact.ump.meal.query
//
// 查询卖家在优惠平台设置的搭配套餐列表，每个套餐包括名称、套餐价格、手淘套餐购买链接
type AlibabaInteractUmpMealQueryAPIResponse struct {
	model.CommonResponse
	AlibabaInteractUmpMealQueryAPIResponseModel
}

// AlibabaInteractUmpMealQueryAPIResponseModel is 淘宝卖家搭配套餐查询 成功返回结果
type AlibabaInteractUmpMealQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_ump_meal_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠平台搭配套餐对象
	Meals []OpenMealDo `json:"meals,omitempty" xml:"meals>open_meal_do,omitempty"`
}
