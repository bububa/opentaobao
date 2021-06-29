package lstspeacker

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstspeacker"
)

/* 
如意音箱音频文件长传 
alibaba.lst.speaker.file.upload

如意音箱音频文件长传
*/
func AlibabaLstSpeakerFileUpload(clt *core.SDKClient, req *lstspeacker.AlibabaLstSpeakerFileUploadRequest, session string) (*lstspeacker.AlibabaLstSpeakerFileUploadAPIResponse, error) {
    var resp lstspeacker.AlibabaLstSpeakerFileUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
