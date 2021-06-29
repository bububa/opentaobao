package smartstore

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/smartstore"
)

/* 
智能硬件互动云货架批量数据回流 
taobao.smartstore.device.iashelf.batch.feedback

智慧门店互动云货架设备批量回流，
只能回流单个设备的批量业务数据
规则：
1. 回流的设备属于当前授权的用户
2. 回流的设备属于当前应用添加
*/
func TaobaoSmartstoreDeviceIashelfBatchFeedback(clt *core.SDKClient, req *smartstore.TaobaoSmartstoreDeviceIashelfBatchFeedbackRequest, session string) (*smartstore.TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIResponse, error) {
    var resp smartstore.TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
