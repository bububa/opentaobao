package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星巴克优惠规则查询 API请求
alibaba.asr.dataservice.promotionrule.query

查询优惠规则，例如星巴克查询优惠规则
*/
type AlibabaAsrDataservicePromotionruleQueryRequest struct {
    model.Params
    // 当前页
    _pageNo   int64
    // 每页数量
    _pageSize   int64
}

// 初始化AlibabaAsrDataservicePromotionruleQueryRequest对象
func NewAlibabaAsrDataservicePromotionruleQueryRequest() *AlibabaAsrDataservicePromotionruleQueryRequest{
    return &AlibabaAsrDataservicePromotionruleQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAsrDataservicePromotionruleQueryRequest) GetApiMethodName() string {
    return "alibaba.asr.dataservice.promotionrule.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAsrDataservicePromotionruleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNo Setter
// 当前页
func (r *AlibabaAsrDataservicePromotionruleQueryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaAsrDataservicePromotionruleQueryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页数量
func (r *AlibabaAsrDataservicePromotionruleQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAsrDataservicePromotionruleQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
