package oversea

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/oversea"
)

/* 
获取文本翻译信息 
alibaba.oversea.translate.get

根据传入的文本信息，获取其目标语言的翻译结果
*/
func AlibabaOverseaTranslateGet(clt *core.SDKClient, req *oversea.AlibabaOverseaTranslateGetRequest, session string) (*oversea.AlibabaOverseaTranslateGetAPIResponse, error) {
    var resp oversea.AlibabaOverseaTranslateGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
