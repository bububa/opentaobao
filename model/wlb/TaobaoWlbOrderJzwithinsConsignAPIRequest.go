package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlborderjzwithinsconsignAPIRequest 家装发货接口 API请求
// taobao.wlb.order.jzwithins.consign
//
// 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口
type TaobaowlborderjzwithinsconsignAPIRequest struct {
	model.Params
	// 淘宝交易订单号
	_tid int64
	// 物流服务商信息
	_tmsPartner *JzPartnerNew
	// 物流服务商信息
	_insPartner *JzPartnerNew
	// 家装物流发货参数
	_jzConsignArgs *JzConsignArgsNew
}

// NewTaobaowlborderjzwithinsconsignRequest 初始化TaobaowlborderjzwithinsconsignAPIRequest对象
func NewTaobaowlborderjzwithinsconsignRequest() *TaobaowlborderjzwithinsconsignAPIRequest {
	return &TaobaowlborderjzwithinsconsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlborderjzwithinsconsignAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.jzwithins.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlborderjzwithinsconsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlborderjzwithinsconsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 淘宝交易订单号
func (r *TaobaowlborderjzwithinsconsignAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaowlborderjzwithinsconsignAPIRequest) GetTid() int64 {
	return r._tid
}

// SetTmsPartner is TmsPartner Setter
// 物流服务商信息
func (r *TaobaowlborderjzwithinsconsignAPIRequest) SetTmsPartner(_tmsPartner *JzPartnerNew) error {
	r._tmsPartner = _tmsPartner
	r.Set("tms_partner", _tmsPartner)
	return nil
}

// GetTmsPartner TmsPartner Getter
func (r TaobaowlborderjzwithinsconsignAPIRequest) GetTmsPartner() *JzPartnerNew {
	return r._tmsPartner
}

// SetInsPartner is InsPartner Setter
// 物流服务商信息
func (r *TaobaowlborderjzwithinsconsignAPIRequest) SetInsPartner(_insPartner *JzPartnerNew) error {
	r._insPartner = _insPartner
	r.Set("ins_partner", _insPartner)
	return nil
}

// GetInsPartner InsPartner Getter
func (r TaobaowlborderjzwithinsconsignAPIRequest) GetInsPartner() *JzPartnerNew {
	return r._insPartner
}

// SetJzConsignArgs is JzConsignArgs Setter
// 家装物流发货参数
func (r *TaobaowlborderjzwithinsconsignAPIRequest) SetJzConsignArgs(_jzConsignArgs *JzConsignArgsNew) error {
	r._jzConsignArgs = _jzConsignArgs
	r.Set("jz_consign_args", _jzConsignArgs)
	return nil
}

// GetJzConsignArgs JzConsignArgs Getter
func (r TaobaowlborderjzwithinsconsignAPIRequest) GetJzConsignArgs() *JzConsignArgsNew {
	return r._jzConsignArgs
}
