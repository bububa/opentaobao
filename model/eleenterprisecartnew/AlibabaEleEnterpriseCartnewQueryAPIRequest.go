package eleenterprisecartnew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseCartnewQueryAPIRequest
新版购物车查询 API请求
alibaba.ele.enterprise.cartnew.query

新版购物车查询 */
type AlibabaEleEnterpriseCartnewQueryAPIRequest struct {
	model.Params
	// 1212
	_phone string
	// 1212
	_latitude string
	// 1212
	_longitude string
	// 餐厅id
	_erestaurantId string
}

// New
