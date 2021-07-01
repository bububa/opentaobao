package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
无线开放视频内容安全检查 API返回值 
taobao.wireless.video.check

无线开放内容检查，提供涉黄暴力政治音视频的异步检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70436.html" target="blank">阿里云内容安全</a>

此API会进行三个场景的审核，检测不通过的视频将被隐藏，用户无法访问被隐藏的视频。

目前，该接口仅支持顽兔空间的视频扫描。
*/
type TaobaoWirelessVideoCheckAPIResponse struct {
    model.CommonResponse
    TaobaoWirelessVideoCheckAPIResponseModel
}

// 无线开放视频内容安全检查 成功返回结果
type TaobaoWirelessVideoCheckAPIResponseModel struct {
    XMLName xml.Name `xml:"wireless_video_check_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 系统自动生成
    Result   *RopResultTo `json:"result,omitempty" xml:"result,omitempty"`
}
