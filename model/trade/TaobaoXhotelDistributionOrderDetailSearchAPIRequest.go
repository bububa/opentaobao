package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhoteldistributionorderdetailsearchAPIRequest 分销渠道订单详情查询 API请求
// taobao.xhotel.distribution.order.detail.search
//
// 该接口用于提供酒店分销渠道的订单详情查询
type TaobaoxhoteldistributionorderdetailsearchAPIRequest struct {
	model.Params
	// 外部分销渠道的订单号（与tid必填其一）
	_distributionOid string
	// 传入用户对应的openId
	_openId string
	// 飞猪的订单号（与distribution_oid必填其一）
	_tid int64
}

// NewTaobaoxhoteldistributionorderdetailsearchRequest 初始化TaobaoxhoteldistributionorderdetailsearchAPIRequest对象
func NewTaobaoxhoteldistributionorderdetailsearchRequest() *TaobaoxhoteldistributionorderdetailsearchAPIRequest {
	return &TaobaoxhoteldistributionorderdetailsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhoteldistributionorderdetailsearchAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.distribution.order.detail.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhoteldistributionorderdetailsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhoteldistributionorderdetailsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributionOid is DistributionOid Setter
// 外部分销渠道的订单号（与tid必填其一）
func (r *TaobaoxhoteldistributionorderdetailsearchAPIRequest) SetDistributionOid(_distributionOid string) error {
	r._distributionOid = _distributionOid
	r.Set("distribution_oid", _distributionOid)
	return nil
}

// GetDistributionOid DistributionOid Getter
func (r TaobaoxhoteldistributionorderdetailsearchAPIRequest) GetDistributionOid() string {
	return r._distributionOid
}

// SetOpenId is OpenId Setter
// 传入用户对应的openId
func (r *TaobaoxhoteldistributionorderdetailsearchAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TaobaoxhoteldistributionorderdetailsearchAPIRequest) GetOpenId() string {
	return r._openId
}

// SetTid is Tid Setter
// 飞猪的订单号（与distribution_oid必填其一）
func (r *TaobaoxhoteldistributionorderdetailsearchAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoxhoteldistributionorderdetailsearchAPIRequest) GetTid() int64 {
	return r._tid
}
