package lstwarehouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstbranddatasharesupplierslistqueryAPIRequest 品牌数据授权的供应商列表 API请求
// alibaba.lst.branddatashare.suppliers.list.query
//
// 品牌商查询品牌数据授权的供应商列表
type AlibabalstbranddatasharesupplierslistqueryAPIRequest struct {
	model.Params
}

// NewAlibabalstbranddatasharesupplierslistqueryRequest 初始化AlibabalstbranddatasharesupplierslistqueryAPIRequest对象
func NewAlibabalstbranddatasharesupplierslistqueryRequest() *AlibabalstbranddatasharesupplierslistqueryAPIRequest {
	return &AlibabalstbranddatasharesupplierslistqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstbranddatasharesupplierslistqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.branddatashare.suppliers.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstbranddatasharesupplierslistqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstbranddatasharesupplierslistqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
