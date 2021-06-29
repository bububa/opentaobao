package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询采购单详情信息 API请求
alibaba.nlife.store.tradedetail.get

根据集团id和采购单号，查询采购单的详情信息
*/
type AlibabaNlifeStoreTradedetailGetRequest struct {
    model.Params
    // 集团采购单号
    procurementNo   string
}

// 初始化AlibabaNlifeStoreTradedetailGetRequest对象
func NewAlibabaNlifeStoreTradedetailGetRequest() *AlibabaNlifeStoreTradedetailGetRequest{
    return &AlibabaNlifeStoreTradedetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreTradedetailGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.tradedetail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreTradedetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProcurementNo Setter
// 集团采购单号
func (r *AlibabaNlifeStoreTradedetailGetRequest) SetProcurementNo(procurementNo string) error {
    r.procurementNo = procurementNo
    r.Set("procurement_no", procurementNo)
    return nil
}

// ProcurementNo Getter
func (r AlibabaNlifeStoreTradedetailGetRequest) GetProcurementNo() string {
    return r.procurementNo
}
