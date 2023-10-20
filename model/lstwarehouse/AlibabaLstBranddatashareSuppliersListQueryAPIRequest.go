package lstwarehouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstBranddatashareSuppliersListQueryAPIRequest 品牌数据授权的供应商列表 API请求
// alibaba.lst.branddatashare.suppliers.list.query
//
// 品牌商查询品牌数据授权的供应商列表
type AlibabaLstBranddatashareSuppliersListQueryAPIRequest struct {
	model.Params
}

// NewAlibabaLstBranddatashareSuppliersListQueryRequest 初始化AlibabaLstBranddatashareSuppliersListQueryAPIRequest对象
func NewAlibabaLstBranddatashareSuppliersListQueryRequest() *AlibabaLstBranddatashareSuppliersListQueryAPIRequest {
	return &AlibabaLstBranddatashareSuppliersListQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstBranddatashareSuppliersListQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstBranddatashareSuppliersListQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.branddatashare.suppliers.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstBranddatashareSuppliersListQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstBranddatashareSuppliersListQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaLstBranddatashareSuppliersListQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstBranddatashareSuppliersListQueryRequest()
	},
}

// GetAlibabaLstBranddatashareSuppliersListQueryRequest 从 sync.Pool 获取 AlibabaLstBranddatashareSuppliersListQueryAPIRequest
func GetAlibabaLstBranddatashareSuppliersListQueryAPIRequest() *AlibabaLstBranddatashareSuppliersListQueryAPIRequest {
	return poolAlibabaLstBranddatashareSuppliersListQueryAPIRequest.Get().(*AlibabaLstBranddatashareSuppliersListQueryAPIRequest)
}

// ReleaseAlibabaLstBranddatashareSuppliersListQueryAPIRequest 将 AlibabaLstBranddatashareSuppliersListQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstBranddatashareSuppliersListQueryAPIRequest(v *AlibabaLstBranddatashareSuppliersListQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstBranddatashareSuppliersListQueryAPIRequest.Put(v)
}
