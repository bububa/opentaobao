package dt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
RT竞价数据回流 API请求
alibaba.nrs.item.rtdata.backflow

回流竞品价格数据，用与后续OCR识别价签数据，做精确化数据纠正
*/
type AlibabaNrsItemRtdataBackflowRequest struct {
    model.Params
    // 入参
    _rtItemPriceTagBackParam   *RtItemPriceTagBackParam
}

// 初始化AlibabaNrsItemRtdataBackflowRequest对象
func NewAlibabaNrsItemRtdataBackflowRequest() *AlibabaNrsItemRtdataBackflowRequest{
    return &AlibabaNrsItemRtdataBackflowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNrsItemRtdataBackflowRequest) GetApiMethodName() string {
    return "alibaba.nrs.item.rtdata.backflow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNrsItemRtdataBackflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RtItemPriceTagBackParam Setter
// 入参
func (r *AlibabaNrsItemRtdataBackflowRequest) SetRtItemPriceTagBackParam(_rtItemPriceTagBackParam *RtItemPriceTagBackParam) error {
    r._rtItemPriceTagBackParam = _rtItemPriceTagBackParam
    r.Set("rt_item_price_tag_back_param", _rtItemPriceTagBackParam)
    return nil
}

// RtItemPriceTagBackParam Getter
func (r AlibabaNrsItemRtdataBackflowRequest) GetRtItemPriceTagBackParam() *RtItemPriceTagBackParam {
    return r._rtItemPriceTagBackParam
}
