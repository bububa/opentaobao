package seaking

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/seaking"
)

/* 
获取图片翻译任务结果 
alibaba.seaking.imagetranslate.result

获取图片翻译任务结果
*/
func AlibabaSeakingImagetranslateResult(clt *core.SDKClient, req *seaking.AlibabaSeakingImagetranslateResultRequest, session string) (*seaking.AlibabaSeakingImagetranslateResultAPIResponse, error) {
    var resp seaking.AlibabaSeakingImagetranslateResultAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
