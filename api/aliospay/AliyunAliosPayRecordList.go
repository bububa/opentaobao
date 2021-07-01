package aliospay

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliospay"
)

/* 
支付记录批量查询接口 
aliyun.alios.pay.record.list

商户用来对账的接口
*/
func AliyunAliosPayRecordList(clt *core.SDKClient, req *aliospay.AliyunAliosPayRecordListAPIRequest, session string) (*aliospay.AliyunAliosPayRecordListAPIResponse, error) {
    var resp aliospay.AliyunAliosPayRecordListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
