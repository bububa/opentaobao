package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaasrdataservicepromotionrulequeryAPIRequest 星巴克优惠规则查询 API请求
// alibaba.asr.dataservice.promotionrule.query
//
// 查询优惠规则，例如星巴克查询优惠规则
type AlibabaasrdataservicepromotionrulequeryAPIRequest struct {
	model.Params
	// 当前页
	_pageNo int64
	// 每页数量
	_pageSize int64
}

// NewAlibabaasrdataservicepromotionrulequeryRequest 初始化AlibabaasrdataservicepromotionrulequeryAPIRequest对象
func NewAlibabaasrdataservicepromotionrulequeryRequest() *AlibabaasrdataservicepromotionrulequeryAPIRequest {
	return &AlibabaasrdataservicepromotionrulequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaasrdataservicepromotionrulequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.asr.dataservice.promotionrule.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaasrdataservicepromotionrulequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaasrdataservicepromotionrulequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNo is PageNo Setter
// 当前页
func (r *AlibabaasrdataservicepromotionrulequeryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AlibabaasrdataservicepromotionrulequeryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页数量
func (r *AlibabaasrdataservicepromotionrulequeryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaasrdataservicepromotionrulequeryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
