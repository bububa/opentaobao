package viapi

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/viapi"
)

/* 
通用文字识别 
aliyun.viapi.ocr.character

获取通用的文字信息。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
func AliyunViapiOcrCharacter(clt *core.SDKClient, req *viapi.AliyunViapiOcrCharacterRequest, session string) (*viapi.AliyunViapiOcrCharacterAPIResponse, error) {
    var resp viapi.AliyunViapiOcrCharacterAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
