package lstwarehouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstbranddatasharestockdataqueryAPIRequest 查询品牌商品实仓库存/周转效能 API请求
// alibaba.lst.branddatashare.stockdata.query
//
// 品牌商查询授权供应商实仓库存数据
type AlibabalstbranddatasharestockdataqueryAPIRequest struct {
	model.Params
	// 入参
	_param *BmSupplierStockDataParam
}

// NewAlibabalstbranddatasharestockdataqueryRequest 初始化AlibabalstbranddatasharestockdataqueryAPIRequest对象
func NewAlibabalstbranddatasharestockdataqueryRequest() *AlibabalstbranddatasharestockdataqueryAPIRequest {
	return &AlibabalstbranddatasharestockdataqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstbranddatasharestockdataqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.branddatashare.stockdata.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstbranddatasharestockdataqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstbranddatasharestockdataqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabalstbranddatasharestockdataqueryAPIRequest) SetParam(_param *BmSupplierStockDataParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabalstbranddatasharestockdataqueryAPIRequest) GetParam() *BmSupplierStockDataParam {
	return r._param
}
