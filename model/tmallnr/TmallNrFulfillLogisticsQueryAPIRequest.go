package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillLogisticsQueryAPIRequest 定时送和极速达配送物流信息查询 API请求
// tmall.nr.fulfill.logistics.query
//
// 发布定时送&amp;极速达物流信息查询服务
type TmallNrFulfillLogisticsQueryAPIRequest struct {
	model.Params
	// 业务标识，dss标识定时送业务；jsd表示极速达业务
	_bizIdentity string
	// 交易主订单号
	_mainOrderId int64
}

// NewTmallNrFulfillLogisticsQueryRequest 初始化TmallNrFulfillLogisticsQueryAPIRequest对象
func NewTmallNrFulfillLogisticsQueryRequest() *TmallNrFulfillLogisticsQueryAPIRequest {
	return &TmallNrFulfillLogisticsQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrFulfillLogisticsQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nr.fulfill.logistics.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrFulfillLogisticsQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrFulfillLogisticsQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizIdentity is BizIdentity Setter
// 业务标识，dss标识定时送业务；jsd表示极速达业务
func (r *TmallNrFulfillLogisticsQueryAPIRequest) SetBizIdentity(_bizIdentity string) error {
	r._bizIdentity = _bizIdentity
	r.Set("biz_identity", _bizIdentity)
	return nil
}

// GetBizIdentity BizIdentity Getter
func (r TmallNrFulfillLogisticsQueryAPIRequest) GetBizIdentity() string {
	return r._bizIdentity
}

// SetMainOrderId is MainOrderId Setter
// 交易主订单号
func (r *TmallNrFulfillLogisticsQueryAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TmallNrFulfillLogisticsQueryAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}
