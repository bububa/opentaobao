package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建经销采购申请 API请求
taobao.fenxiao.dealer.requisitionorder.create

创建经销采购申请
*/
type TaobaoFenxiaoDealerRequisitionorderCreateRequest struct {
    model.Params
    // 配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货
    _logisticsType   string
    // 采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价）
    _orderDetail   []string
    // 收货人所在省份
    _province   string
    // 收货人所在市
    _city   string
    // 收货人所在区
    _district   string
    // 收货人所在街道地址
    _address   string
    // 收货人所在地区邮政编码
    _postCode   string
    // 买家联系电话（此字段和mobile字段至少填写一个）
    _phone   string
    // 买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话）
    _mobile   string
    // 买家姓名（自提方式填写提货人姓名）
    _buyerName   string
    // 身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份）
    _idCardNumber   string
}

// 初始化TaobaoFenxiaoDealerRequisitionorderCreateRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderCreateRequest() *TaobaoFenxiaoDealerRequisitionorderCreateRequest{
    return &TaobaoFenxiaoDealerRequisitionorderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetApiMethodName() string {
    return "taobao.fenxiao.dealer.requisitionorder.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LogisticsType Setter
// 配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货
func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetLogisticsType(_logisticsType string) error {
    r._logisticsType = _logisticsType
    r.Set("logistics_type", _logisticsType)
    return nil
}

// LogisticsType Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetLogisticsType() string {
    return r._logisticsType
}
// OrderDetail Setter
// 采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价）
func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetOrderDetail(_orderDetail []string) error {
    r._orderDetail = _orderDetail
    r.Set("order_detail", _orderDetail)
    return nil
}

// OrderDetail Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetOrderDetail() []string {
    return r._orderDetail
}
// Province Setter
// 收货人所在省份
func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetProvince(_province string) error {
    r._province = _province
    r.Set("province", _province)
    return nil
}

// Province Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetProvince() string {
    return r._province
}
// City Setter
// 收货人所在市
func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetCity() string {
    return r._city
}
// District Setter
// 收货人所在区
func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetDistrict(_district string) error {
    r._district = _district
    r.Set("district", _district)
    return nil
}

// District Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetDistrict() string {
    return r._district
}
// Address Setter
// 收货人所在街道地址
func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetAddress() string {
    return r._address
}
// PostCode Setter
// 收货人所在地区邮政编码
func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetPostCode(_postCode string) error {
    r._postCode = _postCode
    r.Set("post_code", _postCode)
    return nil
}

// PostCode Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetPostCode() string {
    return r._postCode
}
// Phone Setter
// 买家联系电话（此字段和mobile字段至少填写一个）
func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetPhone() string {
    return r._phone
}
// Mobile Setter
// 买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话）
func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetMobile() string {
    return r._mobile
}
// BuyerName Setter
// 买家姓名（自提方式填写提货人姓名）
func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetBuyerName(_buyerName string) error {
    r._buyerName = _buyerName
    r.Set("buyer_name", _buyerName)
    return nil
}

// BuyerName Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetBuyerName() string {
    return r._buyerName
}
// IdCardNumber Setter
// 身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份）
func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetIdCardNumber(_idCardNumber string) error {
    r._idCardNumber = _idCardNumber
    r.Set("id_card_number", _idCardNumber)
    return nil
}

// IdCardNumber Getter
func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetIdCardNumber() string {
    return r._idCardNumber
}
