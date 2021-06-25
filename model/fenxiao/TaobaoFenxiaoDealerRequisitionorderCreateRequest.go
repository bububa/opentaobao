package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建经销采购申请 APIRequest
taobao.fenxiao.dealer.requisitionorder.create

创建经销采购申请
*/
type TaobaoFenxiaoDealerRequisitionorderCreateRequest struct {
    model.Params

    // 配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货
    logisticsType   string 

    // 采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价）
    orderDetail   []String 

    // 收货人所在省份
    province   string 

    // 收货人所在市
    city   string 

    // 收货人所在区
    district   string 

    // 收货人所在街道地址
    address   string 

    // 收货人所在地区邮政编码
    postCode   string 

    // 买家联系电话（此字段和mobile字段至少填写一个）
    phone   string 

    // 买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话）
    mobile   string 

    // 买家姓名（自提方式填写提货人姓名）
    buyerName   string 

    // 身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份）
    idCardNumber   string 

}

func NewTaobaoFenxiaoDealerRequisitionorderCreateRequest() *TaobaoFenxiaoDealerRequisitionorderCreateRequest{
    return &TaobaoFenxiaoDealerRequisitionorderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetApiMethodName() string {
    return "taobao.fenxiao.dealer.requisitionorder.create"
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetLogisticsType(logisticsType string) error {
    r.logisticsType = logisticsType
    r.Set("logistics_type", logisticsType)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetLogisticsType() string {
    return r.logisticsType
}

func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetOrderDetail(orderDetail []String) error {
    r.orderDetail = orderDetail
    r.Set("order_detail", orderDetail)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetOrderDetail() []String {
    return r.orderDetail
}

func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetProvince() string {
    return r.province
}

func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetCity() string {
    return r.city
}

func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetDistrict(district string) error {
    r.district = district
    r.Set("district", district)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetDistrict() string {
    return r.district
}

func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetAddress() string {
    return r.address
}

func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetPostCode(postCode string) error {
    r.postCode = postCode
    r.Set("post_code", postCode)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetPostCode() string {
    return r.postCode
}

func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetPhone() string {
    return r.phone
}

func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetMobile() string {
    return r.mobile
}

func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetBuyerName(buyerName string) error {
    r.buyerName = buyerName
    r.Set("buyer_name", buyerName)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetBuyerName() string {
    return r.buyerName
}

func (r *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetIdCardNumber(idCardNumber string) error {
    r.idCardNumber = idCardNumber
    r.Set("id_card_number", idCardNumber)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderCreateRequest) GetIdCardNumber() string {
    return r.idCardNumber
}

