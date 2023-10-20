package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkordergetAPIRequest 交易订单详情查询 API请求
// alibaba.wdk.order.get
//
// 五道口三江单据查询接口
type AlibabawdkordergetAPIRequest struct {
	model.Params
	// 系统自动生成
	_idListQueryReq *IdListQueryRequest
}

// NewAlibabawdkordergetRequest 初始化AlibabawdkordergetAPIRequest对象
func NewAlibabawdkordergetRequest() *AlibabawdkordergetAPIRequest {
	return &AlibabawdkordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdListQueryReq is IdListQueryReq Setter
// 系统自动生成
func (r *AlibabawdkordergetAPIRequest) SetIdListQueryReq(_idListQueryReq *IdListQueryRequest) error {
	r._idListQueryReq = _idListQueryReq
	r.Set("id_list_query_req", _idListQueryReq)
	return nil
}

// GetIdListQueryReq IdListQueryReq Getter
func (r AlibabawdkordergetAPIRequest) GetIdListQueryReq() *IdListQueryRequest {
	return r._idListQueryReq
}
