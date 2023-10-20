package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelDistributionOrderDetailSearchAPIRequest 分销渠道订单详情查询 API请求
// taobao.xhotel.distribution.order.detail.search
//
// 该接口用于提供酒店分销渠道的订单详情查询
type TaobaoXhotelDistributionOrderDetailSearchAPIRequest struct {
	model.Params
	// 外部分销渠道的订单号（与tid必填其一）
	_distributionOid string
	// 传入用户对应的openId
	_openId string
	// 飞猪的订单号（与distribution_oid必填其一）
	_tid int64
}

// NewTaobaoXhotelDistributionOrderDetailSearchRequest 初始化TaobaoXhotelDistributionOrderDetailSearchAPIRequest对象
func NewTaobaoXhotelDistributionOrderDetailSearchRequest() *TaobaoXhotelDistributionOrderDetailSearchAPIRequest {
	return &TaobaoXhotelDistributionOrderDetailSearchAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelDistributionOrderDetailSearchAPIRequest) Reset() {
	r._distributionOid = ""
	r._openId = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDistributionOrderDetailSearchAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.distribution.order.detail.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDistributionOrderDetailSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelDistributionOrderDetailSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributionOid is DistributionOid Setter
// 外部分销渠道的订单号（与tid必填其一）
func (r *TaobaoXhotelDistributionOrderDetailSearchAPIRequest) SetDistributionOid(_distributionOid string) error {
	r._distributionOid = _distributionOid
	r.Set("distribution_oid", _distributionOid)
	return nil
}

// GetDistributionOid DistributionOid Getter
func (r TaobaoXhotelDistributionOrderDetailSearchAPIRequest) GetDistributionOid() string {
	return r._distributionOid
}

// SetOpenId is OpenId Setter
// 传入用户对应的openId
func (r *TaobaoXhotelDistributionOrderDetailSearchAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TaobaoXhotelDistributionOrderDetailSearchAPIRequest) GetOpenId() string {
	return r._openId
}

// SetTid is Tid Setter
// 飞猪的订单号（与distribution_oid必填其一）
func (r *TaobaoXhotelDistributionOrderDetailSearchAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoXhotelDistributionOrderDetailSearchAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoXhotelDistributionOrderDetailSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelDistributionOrderDetailSearchRequest()
	},
}

// GetTaobaoXhotelDistributionOrderDetailSearchRequest 从 sync.Pool 获取 TaobaoXhotelDistributionOrderDetailSearchAPIRequest
func GetTaobaoXhotelDistributionOrderDetailSearchAPIRequest() *TaobaoXhotelDistributionOrderDetailSearchAPIRequest {
	return poolTaobaoXhotelDistributionOrderDetailSearchAPIRequest.Get().(*TaobaoXhotelDistributionOrderDetailSearchAPIRequest)
}

// ReleaseTaobaoXhotelDistributionOrderDetailSearchAPIRequest 将 TaobaoXhotelDistributionOrderDetailSearchAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelDistributionOrderDetailSearchAPIRequest(v *TaobaoXhotelDistributionOrderDetailSearchAPIRequest) {
	v.Reset()
	poolTaobaoXhotelDistributionOrderDetailSearchAPIRequest.Put(v)
}
