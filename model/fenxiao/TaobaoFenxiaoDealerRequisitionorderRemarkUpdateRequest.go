package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改经销采购单备注 APIRequest
taobao.fenxiao.dealer.requisitionorder.remark.update

供应商修改经销采购单备注
*/
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest struct {
    model.Params

    // 经销采购单ID
    dealerOrderId   int64 

    // 备注留言，可为空
    supplierMemo   string 

    // 旗子的标记，必选。<br/>1-5之间的数字。<br/>非1-5之间，都采用1作为默认。<br/>1:红色<br/>2:黄色<br/>3:绿色<br/>4:蓝色<br/>5:粉红色
    supplierMemoFlag   int64 

}

func NewTaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest() *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest{
    return &TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) GetApiMethodName() string {
    return "taobao.fenxiao.dealer.requisitionorder.remark.update"
}

func (r TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) SetDealerOrderId(dealerOrderId int64) error {
    r.dealerOrderId = dealerOrderId
    r.Set("dealer_order_id", dealerOrderId)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) GetDealerOrderId() int64 {
    return r.dealerOrderId
}

func (r *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) SetSupplierMemo(supplierMemo string) error {
    r.supplierMemo = supplierMemo
    r.Set("supplier_memo", supplierMemo)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) GetSupplierMemo() string {
    return r.supplierMemo
}

func (r *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) SetSupplierMemoFlag(supplierMemoFlag int64) error {
    r.supplierMemoFlag = supplierMemoFlag
    r.Set("supplier_memo_flag", supplierMemoFlag)
    return nil
}

func (r TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) GetSupplierMemoFlag() int64 {
    return r.supplierMemoFlag
}

