package eleenterprisecartnew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseCartnewSaveAPIRequest
新版创建购物车 API请求
alibaba.ele.enterprise.cartnew.save

新版创建购物车 */
type AlibabaEleEnterpriseCartnewSaveAPIRequest struct {
	model.Params
	// 用户11位手机号
	_phone string
	// 用户所在纬度
	_latitude string
	// [[{\"id\": 1526467414,\"new_specs\": [{\"name\": \"规格\",\"value\": \"那么大鲜柠特饮(雪碧) 660ml\"}],\"attrs\": [{\"name\": \"可选小食\",\"value\": \"金黄脆薯格\"}],\"quantity\": 2}]]
	_food string
	// 用户所在经度
	_longitude string
	// 餐厅id
	_erestaurantId string
}

// New
