package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest
创建经销采购申请 API请求
taobao.fenxiao.dealer.requisitionorder.create

创建经销采购申请 */
type TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest struct {
	model.Params
	// 配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货
	_logisticsType string
	// 采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价）
	_orderDetail []string
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

// NewTaobaoFenxiaoDealerRequisitionorderCreateRequest 初始化TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderCreateRequest() *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest {
	return &TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LogisticsType Setter
// 配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货
func (r *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) SetLogisticsType(_logisticsType string) error {
	r._logisticsType = _logisticsType
	r.Set("logistics_type", _logisticsType)
	return nil
}

// Get LogisticsType Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetLogisticsType() string {
	return r._logisticsType
}

// Set is OrderDetail Setter
// 采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价）
func (r *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) SetOrderDetail(_orderDetail []string) error {
	r._orderDetail = _orderDetail
	r.Set("order_detail", _orderDetail)
	return nil
}

// Get OrderDetail Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetOrderDetail() []string {
	return r._orderDetail
}

// Set is Province Setter
// 收货人所在省份
func (r *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// Get Province Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetProvince() string {
	return r._province
}

// Set is City Setter
// 收货人所在市
func (r *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// Get City Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetCity() string {
	return r._city
}

// Set is District Setter
// 收货人所在区
func (r *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) SetDistrict(_district string) error {
	r._district = _district
	r.Set("district", _district)
	return nil
}

// Get District Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetDistrict() string {
	return r._district
}

// Set is Address Setter
// 收货人所在街道地址
func (r *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// Get Address Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetAddress() string {
	return r._address
}

// Set is PostCode Setter
// 收货人所在地区邮政编码
func (r *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) SetPostCode(_postCode string) error {
	r._postCode = _postCode
	r.Set("post_code", _postCode)
	return nil
}

// Get PostCode Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetPostCode() string {
	return r._postCode
}

// Set is Phone Setter
// 买家联系电话（此字段和mobile字段至少填写一个）
func (r *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// Get Phone Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetPhone() string {
	return r._phone
}

// Set is Mobile Setter
// 买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话）
func (r *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// Get Mobile Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetMobile() string {
	return r._mobile
}

// Set is BuyerName Setter
// 买家姓名（自提方式填写提货人姓名）
func (r *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) SetBuyerName(_buyerName string) error {
	r._buyerName = _buyerName
	r.Set("buyer_name", _buyerName)
	return nil
}

// Get BuyerName Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetBuyerName() string {
	return r._buyerName
}

// Set is IdCardNumber Setter
// 身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份）
func (r *TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) SetIdCardNumber(_idCardNumber string) error {
	r._idCardNumber = _idCardNumber
	r.Set("id_card_number", _idCardNumber)
	return nil
}

// Get IdCardNumber Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest) GetIdCardNumber() string {
	return r._idCardNumber
}
