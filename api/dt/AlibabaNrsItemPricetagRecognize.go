package dt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

// AlibabaNrsItemPricetagRecognize 价签识别
// alibaba.nrs.item.pricetag.recognize
//
// 商品价签识别，用于识别RT上传的竞品分析照片，返回价签内容
func AlibabaNrsItemPricetagRecognize(clt *core.SDKClient, req *dt.AlibabaNrsItemPricetagRecognizeAPIRequest, resp *dt.AlibabaNrsItemPricetagRecognizeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
