package dt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/dt"
)

/* 
价签识别 
alibaba.nrs.item.pricetag.recognize

商品价签识别，用于识别RT上传的竞品分析照片，返回价签内容
*/
func AlibabaNrsItemPricetagRecognize(clt *core.SDKClient, req *dt.AlibabaNrsItemPricetagRecognizeRequest, session string) (*dt.AlibabaNrsItemPricetagRecognizeAPIResponse, error) {
    var resp dt.AlibabaNrsItemPricetagRecognizeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
