package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询采购单详情信息 APIRequest
alibaba.nlife.store.tradedetail.get

根据集团id和采购单号，查询采购单的详情信息
*/
type AlibabaNlifeStoreTradedetailGetRequest struct {
    model.Params

    // 集团采购单号
    procurementNo   string 

}

func NewAlibabaNlifeStoreTradedetailGetRequest() *AlibabaNlifeStoreTradedetailGetRequest{
    return &AlibabaNlifeStoreTradedetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNlifeStoreTradedetailGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.tradedetail.get"
}

func (r AlibabaNlifeStoreTradedetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNlifeStoreTradedetailGetRequest) SetProcurementNo(procurementNo string) error {
    r.procurementNo = procurementNo
    r.Set("procurement_no", procurementNo)
    return nil
}

func (r AlibabaNlifeStoreTradedetailGetRequest) GetProcurementNo() string {
    return r.procurementNo
}

