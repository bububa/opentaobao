package lstwarehouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstBranddatashareStockdataQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstBranddatashareStockdataQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.branddatashare.stockdata.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstBranddatashareStockdataQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstBranddatashareStockdataQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaLstBranddatashareStockdataQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstBranddatashareStockdataQueryRequest()
	},
}

// GetAlibabaLstBranddatashareStockdataQueryRequest 从 sync.Pool 获取 AlibabaLstBranddatashareStockdataQueryAPIRequest
func GetAlibabaLstBranddatashareStockdataQueryAPIRequest() *AlibabaLstBranddatashareStockdataQueryAPIRequest {
	return poolAlibabaLstBranddatashareStockdataQueryAPIRequest.Get().(*AlibabaLstBranddatashareStockdataQueryAPIRequest)
}

// ReleaseAlibabaLstBranddatashareStockdataQueryAPIRequest 将 AlibabaLstBranddatashareStockdataQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstBranddatashareStockdataQueryAPIRequest(v *AlibabaLstBranddatashareStockdataQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstBranddatashareStockdataQueryAPIRequest.Put(v)
}
