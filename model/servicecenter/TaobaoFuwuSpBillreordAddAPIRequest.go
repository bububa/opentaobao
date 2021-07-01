package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
内购服务确认单明细上传接口 API请求
taobao.fuwu.sp.billreord.add

isv能通过该接口上传确认单明细数据
*/
type TaobaoFuwuSpBillreordAddAPIRequest struct {
    model.Params
    // 确认单的账单明细
    _paramBillRecordDTO   *BillRecordDto
}

// 初始化TaobaoFuwuSpBillreordAddAPIRequest对象
func NewTaobaoFuwuSpBillreordAddRequest() *TaobaoFuwuSpBillreordAddAPIRequest{
    return &TaobaoFuwuSpBillreordAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFuwuSpBillreordAddAPIRequest) GetApiMethodName() string {
    return "taobao.fuwu.sp.billreord.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFuwuSpBillreordAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBillRecordDTO Setter
// 确认单的账单明细
func (r *TaobaoFuwuSpBillreordAddAPIRequest) SetParamBillRecordDTO(_paramBillRecordDTO *BillRecordDto) error {
    r._paramBillRecordDTO = _paramBillRecordDTO
    r.Set("param_bill_record_d_t_o", _paramBillRecordDTO)
    return nil
}

// ParamBillRecordDTO Getter
func (r TaobaoFuwuSpBillreordAddAPIRequest) GetParamBillRecordDTO() *BillRecordDto {
    return r._paramBillRecordDTO
}
