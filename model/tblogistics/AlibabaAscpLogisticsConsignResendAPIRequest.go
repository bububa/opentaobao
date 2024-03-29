package tblogistics

import (
	"net/url"
	"sync"

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
	// feature参数格式，KV之间用“=”分隔，多个key之间用”;”分隔 ，范例: instantMobilePhoneNumber=12345678910表示同城配送物流公司的物流订单收货人手机号，支持11位真实号和15位隐私号"12345678910-1234"
	_feature string
}

// NewAlibabaAscpLogisticsConsignResendRequest 初始化AlibabaAscpLogisticsConsignResendAPIRequest对象
func NewAlibabaAscpLogisticsConsignResendRequest() *AlibabaAscpLogisticsConsignResendAPIRequest {
	return &AlibabaAscpLogisticsConsignResendAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpLogisticsConsignResendAPIRequest) Reset() {
	r._consignPkgs = r._consignPkgs[:0]
	r._tid = ""
	r._subTids = ""
	r._feature = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsConsignResendAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.consign.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsConsignResendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsConsignResendAPIRequest) GetRawParams() model.Params {
	return r.Params
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

// SetFeature is Feature Setter
// feature参数格式，KV之间用“=”分隔，多个key之间用”;”分隔 ，范例: instantMobilePhoneNumber=12345678910表示同城配送物流公司的物流订单收货人手机号，支持11位真实号和15位隐私号&#34;12345678910-1234&#34;
func (r *AlibabaAscpLogisticsConsignResendAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r AlibabaAscpLogisticsConsignResendAPIRequest) GetFeature() string {
	return r._feature
}

var poolAlibabaAscpLogisticsConsignResendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpLogisticsConsignResendRequest()
	},
}

// GetAlibabaAscpLogisticsConsignResendRequest 从 sync.Pool 获取 AlibabaAscpLogisticsConsignResendAPIRequest
func GetAlibabaAscpLogisticsConsignResendAPIRequest() *AlibabaAscpLogisticsConsignResendAPIRequest {
	return poolAlibabaAscpLogisticsConsignResendAPIRequest.Get().(*AlibabaAscpLogisticsConsignResendAPIRequest)
}

// ReleaseAlibabaAscpLogisticsConsignResendAPIRequest 将 AlibabaAscpLogisticsConsignResendAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpLogisticsConsignResendAPIRequest(v *AlibabaAscpLogisticsConsignResendAPIRequest) {
	v.Reset()
	poolAlibabaAscpLogisticsConsignResendAPIRequest.Put(v)
}
