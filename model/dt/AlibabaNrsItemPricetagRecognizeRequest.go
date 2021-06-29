package dt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价签识别 API请求
alibaba.nrs.item.pricetag.recognize

商品价签识别，用于识别RT上传的竞品分析照片，返回价签内容
*/
type AlibabaNrsItemPricetagRecognizeRequest struct {
    model.Params
    // 图片数据
    _imgByteArr   *model.File
    // 价签识别参数
    _priceTagParam   *PriceTagReqParam
}

// 初始化AlibabaNrsItemPricetagRecognizeRequest对象
func NewAlibabaNrsItemPricetagRecognizeRequest() *AlibabaNrsItemPricetagRecognizeRequest{
    return &AlibabaNrsItemPricetagRecognizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNrsItemPricetagRecognizeRequest) GetApiMethodName() string {
    return "alibaba.nrs.item.pricetag.recognize"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNrsItemPricetagRecognizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImgByteArr Setter
// 图片数据
func (r *AlibabaNrsItemPricetagRecognizeRequest) SetImgByteArr(_imgByteArr *model.File) error {
    r._imgByteArr = _imgByteArr
    r.Set("img_byte_arr", _imgByteArr)
    return nil
}

// ImgByteArr Getter
func (r AlibabaNrsItemPricetagRecognizeRequest) GetImgByteArr() *model.File {
    return r._imgByteArr
}
// PriceTagParam Setter
// 价签识别参数
func (r *AlibabaNrsItemPricetagRecognizeRequest) SetPriceTagParam(_priceTagParam *PriceTagReqParam) error {
    r._priceTagParam = _priceTagParam
    r.Set("price_tag_param", _priceTagParam)
    return nil
}

// PriceTagParam Getter
func (r AlibabaNrsItemPricetagRecognizeRequest) GetPriceTagParam() *PriceTagReqParam {
    return r._priceTagParam
}
