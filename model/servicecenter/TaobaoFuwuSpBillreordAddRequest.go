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
type TaobaoFuwuSpBillreordAddRequest struct {
    model.Params
    // 确认单的账单明细
    paramBillRecordDTO   *BillRecordDto
}

// 初始化TaobaoFuwuSpBillreordAddRequest对象
func NewTaobaoFuwuSpBillreordAddRequest() *TaobaoFuwuSpBillreordAddRequest{
    return &TaobaoFuwuSpBillreordAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFuwuSpBillreordAddRequest) GetApiMethodName() string {
    return "taobao.fuwu.sp.billreord.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFuwuSpBillreordAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBillRecordDTO Setter
// 确认单的账单明细
func (r *TaobaoFuwuSpBillreordAddRequest) SetParamBillRecordDTO(paramBillRecordDTO *BillRecordDto) error {
    r.paramBillRecordDTO = paramBillRecordDTO
    r.Set("param_bill_record_d_t_o", paramBillRecordDTO)
    return nil
}

// ParamBillRecordDTO Getter
func (r TaobaoFuwuSpBillreordAddRequest) GetParamBillRecordDTO() *BillRecordDto {
    return r.paramBillRecordDTO
}
