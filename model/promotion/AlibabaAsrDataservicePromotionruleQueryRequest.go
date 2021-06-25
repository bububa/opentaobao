package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星巴克优惠规则查询 APIRequest
alibaba.asr.dataservice.promotionrule.query

查询优惠规则，例如星巴克查询优惠规则
*/
type AlibabaAsrDataservicePromotionruleQueryRequest struct {
    model.Params

    // 当前页
    pageNo   int64 

    // 每页数量
    pageSize   int64 

}

func NewAlibabaAsrDataservicePromotionruleQueryRequest() *AlibabaAsrDataservicePromotionruleQueryRequest{
    return &AlibabaAsrDataservicePromotionruleQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAsrDataservicePromotionruleQueryRequest) GetApiMethodName() string {
    return "alibaba.asr.dataservice.promotionrule.query"
}

func (r AlibabaAsrDataservicePromotionruleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAsrDataservicePromotionruleQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AlibabaAsrDataservicePromotionruleQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *AlibabaAsrDataservicePromotionruleQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAsrDataservicePromotionruleQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

