package alitrippoi

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitrippoi"
)

/* 
穷游内容写入接口 
alitrip.platform.content.raw.add

穷游内容写入飞猪接口
*/
func AlitripPlatformContentRawAdd(clt *core.SDKClient, req *alitrippoi.AlitripPlatformContentRawAddAPIRequest, session string) (*alitrippoi.AlitripPlatformContentRawAddAPIResponse, error) {
    var resp alitrippoi.AlitripPlatformContentRawAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
