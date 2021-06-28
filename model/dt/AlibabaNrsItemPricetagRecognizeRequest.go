package dt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价签识别 APIRequest
alibaba.nrs.item.pricetag.recognize

商品价签识别，用于识别RT上传的竞品分析照片，返回价签内容
*/
type AlibabaNrsItemPricetagRecognizeRequest struct {
    model.Params

    // 图片数据
    imgByteArr   []*model.File 

    // 价签识别参数
    priceTagParam   *PriceTagReqParam 

}

func NewAlibabaNrsItemPricetagRecognizeRequest() *AlibabaNrsItemPricetagRecognizeRequest{
    return &AlibabaNrsItemPricetagRecognizeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNrsItemPricetagRecognizeRequest) GetApiMethodName() string {
    return "alibaba.nrs.item.pricetag.recognize"
}

func (r AlibabaNrsItemPricetagRecognizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNrsItemPricetagRecognizeRequest) SetImgByteArr(imgByteArr []*model.File) error {
    r.imgByteArr = imgByteArr
    r.Set("img_byte_arr", imgByteArr)
    return nil
}

func (r AlibabaNrsItemPricetagRecognizeRequest) GetImgByteArr() []*model.File {
    return r.imgByteArr
}

func (r *AlibabaNrsItemPricetagRecognizeRequest) SetPriceTagParam(priceTagParam *PriceTagReqParam) error {
    r.priceTagParam = priceTagParam
    r.Set("price_tag_param", priceTagParam)
    return nil
}

func (r AlibabaNrsItemPricetagRecognizeRequest) GetPriceTagParam() *PriceTagReqParam {
    return r.priceTagParam
}

