package degoperation

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/degoperation"
)

/* 
中奖生成二维码 
taobao.degoperation.createqrcode

用户中奖后，生成二维码图片链接
*/
func TaobaoDegoperationCreateqrcode(clt *core.SDKClient, req *degoperation.TaobaoDegoperationCreateqrcodeRequest, session string) (*degoperation.TaobaoDegoperationCreateqrcodeAPIResponse, error) {
    var resp degoperation.TaobaoDegoperationCreateqrcodeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
