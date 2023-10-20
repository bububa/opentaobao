package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaifpfulfillwarehousetokenqueryAPIRequest 同城令牌打印接口 API请求
// alibaba.ifp.fulfill.warehouse.token.query
//
// 用于仓内作业打印包裹信息
type AlibabaifpfulfillwarehousetokenqueryAPIRequest struct {
	model.Params
	// 入参
	_packageQueryDTO *PackageQueryDto
}

// NewAlibabaifpfulfillwarehousetokenqueryRequest 初始化AlibabaifpfulfillwarehousetokenqueryAPIRequest对象
func NewAlibabaifpfulfillwarehousetokenqueryRequest() *AlibabaifpfulfillwarehousetokenqueryAPIRequest {
	return &AlibabaifpfulfillwarehousetokenqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaifpfulfillwarehousetokenqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ifp.fulfill.warehouse.token.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaifpfulfillwarehousetokenqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaifpfulfillwarehousetokenqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPackageQueryDTO is PackageQueryDTO Setter
// 入参
func (r *AlibabaifpfulfillwarehousetokenqueryAPIRequest) SetPackageQueryDTO(_packageQueryDTO *PackageQueryDto) error {
	r._packageQueryDTO = _packageQueryDTO
	r.Set("package_query_d_t_o", _packageQueryDTO)
	return nil
}

// GetPackageQueryDTO PackageQueryDTO Getter
func (r AlibabaifpfulfillwarehousetokenqueryAPIRequest) GetPackageQueryDTO() *PackageQueryDto {
	return r._packageQueryDTO
}
