package lstwarehouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstBranddatashareStockdataQueryAPIRequest 查询品牌商品实仓库存/周转效能 API请求
// alibaba.lst.branddatashare.stockdata.query
//
// 品牌商查询授权供应商实仓库存数据
type AlibabaLstBranddatashareStockdataQueryAPIRequest struct {
	model.Params
	// 入参
	_param *BmSupplierStockDataParam
}

// NewAlibabaLstBranddatashareStockdataQueryRequest 初始化AlibabaLstBranddatashareStockdataQueryAPIRequest对象
func NewAlibabaLstBranddatashareStockdataQueryRequest() *AlibabaLstBranddatashareStockdataQueryAPIRequest {
	return &AlibabaLstBranddatashareStockdataQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstBranddatashareStockdataQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.branddatashare.stockdata.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstBranddatashareStockdataQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaLstBranddatashareStockdataQueryAPIRequest) SetParam(_param *BmSupplierStockDataParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaLstBranddatashareStockdataQueryAPIRequest) GetParam() *BmSupplierStockDataParam {
	return r._param
}
