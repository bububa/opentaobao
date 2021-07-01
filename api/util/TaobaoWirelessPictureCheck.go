package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
无线开放图片内容安全检查 
taobao.wireless.picture.check

无线开放内容检查，提供涉黄暴力政治图片检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70292.html" target="blank">阿里云内容安全</a>
此API会进行两个场景审核，平均RT为1s。
*/
func TaobaoWirelessPictureCheck(clt *core.SDKClient, req *util.TaobaoWirelessPictureCheckAPIRequest, session string) (*util.TaobaoWirelessPictureCheckAPIResponse, error) {
    var resp util.TaobaoWirelessPictureCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
