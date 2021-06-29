package damaiticklet

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/damaiticklet"
)

/* 
票夹-动态二维码-解码 
alibaba.damai.ticklet.qrcode.decode

对于票夹的动态二维码进行解码
*/
func AlibabaDamaiTickletQrcodeDecode(clt *core.SDKClient, req *damaiticklet.AlibabaDamaiTickletQrcodeDecodeRequest, session string) (*damaiticklet.AlibabaDamaiTickletQrcodeDecodeAPIResponse, error) {
    var resp damaiticklet.AlibabaDamaiTickletQrcodeDecodeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
