package wangwang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniukefuevalgetAPIRequest 客服评价详情接口 API请求
// taobao.qianniu.kefueval.get
//
// 获取买家对客服的服务评价
type TaobaoqianniukefuevalgetAPIRequest struct {
	model.Params
	// 客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick
	_queryIds string
	// 开始时间，格式yyyyMMddHHmmss
	_btime string
	// 结束时间，格式yyyyMMddHHmmss
	_etime string
}

// NewTaobaoqianniukefuevalgetRequest 初始化TaobaoqianniukefuevalgetAPIRequest对象
func NewTaobaoqianniukefuevalgetRequest() *TaobaoqianniukefuevalgetAPIRequest {
	return &TaobaoqianniukefuevalgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqianniukefuevalgetAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.kefueval.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqianniukefuevalgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqianniukefuevalgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryIds is QueryIds Setter
// 客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick
func (r *TaobaoqianniukefuevalgetAPIRequest) SetQueryIds(_queryIds string) error {
	r._queryIds = _queryIds
	r.Set("query_ids", _queryIds)
	return nil
}

// GetQueryIds QueryIds Getter
func (r TaobaoqianniukefuevalgetAPIRequest) GetQueryIds() string {
	return r._queryIds
}

// SetBtime is Btime Setter
// 开始时间，格式yyyyMMddHHmmss
func (r *TaobaoqianniukefuevalgetAPIRequest) SetBtime(_btime string) error {
	r._btime = _btime
	r.Set("btime", _btime)
	return nil
}

// GetBtime Btime Getter
func (r TaobaoqianniukefuevalgetAPIRequest) GetBtime() string {
	return r._btime
}

// SetEtime is Etime Setter
// 结束时间，格式yyyyMMddHHmmss
func (r *TaobaoqianniukefuevalgetAPIRequest) SetEtime(_etime string) error {
	r._etime = _etime
	r.Set("etime", _etime)
	return nil
}

// GetEtime Etime Getter
func (r TaobaoqianniukefuevalgetAPIRequest) GetEtime() string {
	return r._etime
}
