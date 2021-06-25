package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
获取手机banner图 
alibaba.xiami.api.mobile.figureimage.get

获取手机banner图
*/
func AlibabaXiamiApiMobileFigureimageGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiMobileFigureimageGetRequest, session string) (*xiami.AlibabaXiamiApiMobileFigureimageGetResponse, error) {
    var resp xiami.AlibabaXiamiApiMobileFigureimageGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
