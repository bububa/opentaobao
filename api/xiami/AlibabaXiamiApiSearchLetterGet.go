package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
搜索接口（首字母） 
alibaba.xiami.api.search.letter.get

搜索接口（首字母）
*/
func AlibabaXiamiApiSearchLetterGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiSearchLetterGetRequest, session string) (*xiami.AlibabaXiamiApiSearchLetterGetResponse, error) {
    var resp xiami.AlibabaXiamiApiSearchLetterGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
