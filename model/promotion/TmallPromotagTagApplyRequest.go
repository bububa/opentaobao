package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠标签申请 APIRequest
tmall.promotag.tag.apply

创建优惠标签
*/
type TmallPromotagTagApplyRequest struct {
    model.Params

    // 标签名称。
注意在UMP中使用新人群标签top变成大写的“NEW_” 如：老标签是top1234，新标签是NEW_1234 。
    tagName   string 

    // 标签用途描述
    tagDesc   string 

    // 标签开始时间
    startTime   string 

    // 标签结束时间
    endTime   string 

}

func NewTmallPromotagTagApplyRequest() *TmallPromotagTagApplyRequest{
    return &TmallPromotagTagApplyRequest{
        Params: model.NewParams(),
    }
}

func (r TmallPromotagTagApplyRequest) GetApiMethodName() string {
    return "tmall.promotag.tag.apply"
}

func (r TmallPromotagTagApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallPromotagTagApplyRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

func (r TmallPromotagTagApplyRequest) GetTagName() string {
    return r.tagName
}

func (r *TmallPromotagTagApplyRequest) SetTagDesc(tagDesc string) error {
    r.tagDesc = tagDesc
    r.Set("tag_desc", tagDesc)
    return nil
}

func (r TmallPromotagTagApplyRequest) GetTagDesc() string {
    return r.tagDesc
}

func (r *TmallPromotagTagApplyRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TmallPromotagTagApplyRequest) GetStartTime() string {
    return r.startTime
}

func (r *TmallPromotagTagApplyRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TmallPromotagTagApplyRequest) GetEndTime() string {
    return r.endTime
}

