package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsConsignResendAPIRequest 修改物流公司和运单号 API请求
// alibaba.ascp.logistics.consign.resend
//
// 支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：
// 1、必须是已发货订单，自己联系发货的必须50天内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；
// 2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司
type AlibabaAscpLogisticsConsignResendAPIRequest struct {
	model.Params
	// 包裹包含的运单号及快递公司编号,多包裹时，需要包含所有包裹的运单号等信息
	_consignPkgs []TopConsignPkgRequest
	// 订单id
	_tid string
	// 拆单子订单列表，对应的数据是：子订单号列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集
	_subTids string
}

// NewAlibabaAscpLogisticsConsignResendRequest 初始化AlibabaAscpLogisticsConsignResendAPIRequest对象
func NewAlibabaAscpLogisticsConsignResendRequest() *AlibabaAscpLogisticsConsignResendAPIRequest {
	return &AlibabaAscpLogisticsConsignResendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsConsignResendAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.consign.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsConsignResendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetConsignPkgs is ConsignPkgs Setter
// 包裹包含的运单号及快递公司编号,多包裹时，需要包含所有包裹的运单号等信息
func (r *AlibabaAscpLogisticsConsignResendAPIRequest) SetConsignPkgs(_consignPkgs []TopConsignPkgRequest) error {
	r._consignPkgs = _consignPkgs
	r.Set("consign_pkgs", _consignPkgs)
	return nil
}

// GetConsignPkgs ConsignPkgs Getter
func (r AlibabaAscpLogisticsConsignResendAPIRequest) GetConsignPkgs() []TopConsignPkgRequest {
	return r._consignPkgs
}

// SetTid is Tid Setter
// 订单id
func (r *AlibabaAscpLogisticsConsignResendAPIRequest) SetTid(_tid string) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r AlibabaAscpLogisticsConsignResendAPIRequest) GetTid() string {
	return r._tid
}

// SetSubTids is SubTids Setter
// 拆单子订单列表，对应的数据是：子订单号列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集
func (r *AlibabaAscpLogisticsConsignResendAPIRequest) SetSubTids(_subTids string) error {
	r._subTids = _subTids
	r.Set("sub_tids", _subTids)
	return nil
}

// GetSubTids SubTids Getter
func (r AlibabaAscpLogisticsConsignResendAPIRequest) GetSubTids() string {
	return r._subTids
}
