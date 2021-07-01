package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
录音文件下载 
alibaba.aliqin.fc.voice.record.geturl

完成录音文件的下载地址获取操作
*/
func AlibabaAliqinFcVoiceRecordGeturl(clt *core.SDKClient, req *alicom.AlibabaAliqinFcVoiceRecordGeturlAPIRequest, session string) (*alicom.AlibabaAliqinFcVoiceRecordGeturlAPIResponse, error) {
    var resp alicom.AlibabaAliqinFcVoiceRecordGeturlAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
