package wenyuvideo

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wenyuvideo"
)

/* 
只看TA 
youku.wenyuvideo.seeta.get

只看Ta对外输出
*/
func YoukuWenyuvideoSeetaGet(clt *core.SDKClient, req *wenyuvideo.YoukuWenyuvideoSeetaGetAPIRequest, session string) (*wenyuvideo.YoukuWenyuvideoSeetaGetAPIResponse, error) {
    var resp wenyuvideo.YoukuWenyuvideoSeetaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
