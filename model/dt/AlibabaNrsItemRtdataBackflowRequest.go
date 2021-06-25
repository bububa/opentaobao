package dt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
RT竞价数据回流 APIRequest
alibaba.nrs.item.rtdata.backflow

回流竞品价格数据，用与后续OCR识别价签数据，做精确化数据纠正
*/
type AlibabaNrsItemRtdataBackflowRequest struct {
    model.Params

    // 入参
    rtItemPriceTagBackParam   *RtItemPriceTagBackParam 

}

func NewAlibabaNrsItemRtdataBackflowRequest() *AlibabaNrsItemRtdataBackflowRequest{
    return &AlibabaNrsItemRtdataBackflowRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNrsItemRtdataBackflowRequest) GetApiMethodName() string {
    return "alibaba.nrs.item.rtdata.backflow"
}

func (r AlibabaNrsItemRtdataBackflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNrsItemRtdataBackflowRequest) SetRtItemPriceTagBackParam(rtItemPriceTagBackParam *RtItemPriceTagBackParam) error {
    r.rtItemPriceTagBackParam = rtItemPriceTagBackParam
    r.Set("rt_item_price_tag_back_param", rtItemPriceTagBackParam)
    return nil
}

func (r AlibabaNrsItemRtdataBackflowRequest) GetRtItemPriceTagBackParam() *RtItemPriceTagBackParam {
    return r.rtItemPriceTagBackParam
}

