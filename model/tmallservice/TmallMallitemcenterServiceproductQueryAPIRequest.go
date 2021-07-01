package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMallitemcenterServiceproductQueryAPIRequest
天猫服务商服务产品查询 API请求
tmall.mallitemcenter.serviceproduct.query

查询天猫服务的服务商发布的服务产品 */
type TmallMallitemcenterServiceproductQueryAPIRequest struct {
	model.Params
	// 服务产品id
	_id int64
	// 产品状态
	_status int64
	// 服务名称
	_serviceCode string
	// 产品类型
	_serviceProductType int64
}

// New
