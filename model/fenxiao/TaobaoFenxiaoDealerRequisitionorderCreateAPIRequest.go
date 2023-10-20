package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaodealerrequisitionordercreateAPIRequest 创建经销采购申请 API请求
// taobao.fenxiao.dealer.requisitionorder.create
//
// 创建经销采购申请
type TaobaofenxiaodealerrequisitionordercreateAPIRequest struct {
	model.Params
	// 采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价）
	_orderDetail []string
	// 配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货
	_logisticsType string
	// 收货人所在省份
	_province string
	// 收货人所在市
	_city string
	// 收货人所在区
	_district string
	// 收货人所在街道地址
	_address string
	// 收货人所在地区邮政编码
	_postCode string
	// 买家联系电话（此字段和mobile字段至少填写一个）
	_phone string
	// 买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话）
	_mobile string
	// 买家姓名（自提方式填写提货人姓名）
	_buyerName string
	// 身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份）
	_idCardNumber string
}

// NewTaobaofenxiaodealerrequisitionordercreateRequest 初始化TaobaofenxiaodealerrequisitionordercreateAPIRequest对象
func NewTaobaofenxiaodealerrequisitionordercreateRequest() *TaobaofenxiaodealerrequisitionordercreateAPIRequest {
	return &TaobaofenxiaodealerrequisitionordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderDetail is OrderDetail Setter
// 采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价）
func (r *TaobaofenxiaodealerrequisitionordercreateAPIRequest) SetOrderDetail(_orderDetail []string) error {
	r._orderDetail = _orderDetail
	r.Set("order_detail", _orderDetail)
	return nil
}

// GetOrderDetail OrderDetail Getter
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetOrderDetail() []string {
	return r._orderDetail
}

// SetLogisticsType is LogisticsType Setter
// 配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货
func (r *TaobaofenxiaodealerrequisitionordercreateAPIRequest) SetLogisticsType(_logisticsType string) error {
	r._logisticsType = _logisticsType
	r.Set("logistics_type", _logisticsType)
	return nil
}

// GetLogisticsType LogisticsType Getter
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetLogisticsType() string {
	return r._logisticsType
}

// SetProvince is Province Setter
// 收货人所在省份
func (r *TaobaofenxiaodealerrequisitionordercreateAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetProvince() string {
	return r._province
}

// SetCity is City Setter
// 收货人所在市
func (r *TaobaofenxiaodealerrequisitionordercreateAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetCity() string {
	return r._city
}

// SetDistrict is District Setter
// 收货人所在区
func (r *TaobaofenxiaodealerrequisitionordercreateAPIRequest) SetDistrict(_district string) error {
	r._district = _district
	r.Set("district", _district)
	return nil
}

// GetDistrict District Getter
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetDistrict() string {
	return r._district
}

// SetAddress is Address Setter
// 收货人所在街道地址
func (r *TaobaofenxiaodealerrequisitionordercreateAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetAddress() string {
	return r._address
}

// SetPostCode is PostCode Setter
// 收货人所在地区邮政编码
func (r *TaobaofenxiaodealerrequisitionordercreateAPIRequest) SetPostCode(_postCode string) error {
	r._postCode = _postCode
	r.Set("post_code", _postCode)
	return nil
}

// GetPostCode PostCode Getter
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetPostCode() string {
	return r._postCode
}

// SetPhone is Phone Setter
// 买家联系电话（此字段和mobile字段至少填写一个）
func (r *TaobaofenxiaodealerrequisitionordercreateAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetPhone() string {
	return r._phone
}

// SetMobile is Mobile Setter
// 买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话）
func (r *TaobaofenxiaodealerrequisitionordercreateAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetMobile() string {
	return r._mobile
}

// SetBuyerName is BuyerName Setter
// 买家姓名（自提方式填写提货人姓名）
func (r *TaobaofenxiaodealerrequisitionordercreateAPIRequest) SetBuyerName(_buyerName string) error {
	r._buyerName = _buyerName
	r.Set("buyer_name", _buyerName)
	return nil
}

// GetBuyerName BuyerName Getter
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetBuyerName() string {
	return r._buyerName
}

// SetIdCardNumber is IdCardNumber Setter
// 身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份）
func (r *TaobaofenxiaodealerrequisitionordercreateAPIRequest) SetIdCardNumber(_idCardNumber string) error {
	r._idCardNumber = _idCardNumber
	r.Set("id_card_number", _idCardNumber)
	return nil
}

// GetIdCardNumber IdCardNumber Getter
func (r TaobaofenxiaodealerrequisitionordercreateAPIRequest) GetIdCardNumber() string {
	return r._idCardNumber
}
