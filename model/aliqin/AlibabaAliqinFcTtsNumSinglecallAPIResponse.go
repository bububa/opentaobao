package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
文本转语音通知 API返回值 
alibaba.aliqin.fc.tts.num.singlecall

向指定手机号码发起单向呼叫，将文本模板内容转化为语音播放给被叫方。使用前需要在阿里大于管理中心添加去电显示号码与文本转语音模板。
*/
type AlibabaAliqinFcTtsNumSinglecallAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcTtsNumSinglecallAPIResponseModel
}

// 文本转语音通知 成功返回结果
type AlibabaAliqinFcTtsNumSinglecallAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_tts_num_singlecall_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回参数
    Result   *AlibabaAliqinFcTtsNumSinglecallBizResult `json:"result,omitempty" xml:"result,omitempty"`
}
